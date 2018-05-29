import json
import os
import sys
from typing import Dict, List, Tuple

package = "obsws"

doc = "https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md"

disclaimer = """\
// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py\
"""


type_map = {
    "bool": "bool",
    "boolean": "bool",
    "int": "int",
    "float": "float64",
    "double": "float64",
    "string": "string",
    "array": "[]string",
    "object": "map[string]interface{}",
    "array of objects": "[]map[string]interface{}",
    "scene": "interface{}",  # String?
    "object|array": "interface{}",
    "scene|array": "interface{}",
    "source|array": "interface{}",
}

unknown_types = [
    "object|array",
    "scene|array",
    "scene",  # String?
    "source|array",
]


def optional_type(s: str) -> Tuple[str, bool]:
    """Determine if a type is optional and parse the actual type name."""
    if s.endswith("(optional)"):
        return s[:s.find("(optional)")].strip(), True
    return s, False


def process_json(d: Dict):
    """Generate Go code for the entire protocol file."""
    ("--events" in sys.argv or "--all" in sys.argv) and gen_events(d["events"])
    ("--requests" in sys.argv or "--all" in sys.argv) and gen_requests(d["requests"])
    ("--typeswitches" in sys.argv or "--all" in sys.argv) and gen_typeswitches(d)


def gen_category(prefix: str, category: str, data: Dict):
    """Generate all events or requests in one category."""
    func = gen_event if prefix == "events" else gen_request
    content = "\n\n".join(func(event) for event in data)
    with open(f"{prefix}_{category}.go".replace(" ", "_"), "w") as f:
        f.write(f"""\
        package {package}

        {disclaimer}

        {content}
        """)


def gen_events(events: Dict):
    """Generate all events."""
    for category, data in events.items():
        gen_category("events", category, data)


def gen_event(data: Dict) -> str:
    """Write Go code with a type definition and interface functions."""
    reserved = ["Type", "StreamTC", "RecTC"]

    if data.get("returns"):
        struct = f"""\
        type {data["name"]}Event struct {{
            {go_struct_variables(go_variables(data["returns"], reserved))}
            _event
        }}\
        """
    else:
        struct = f"type {data['name']}Event _event"

    description = newlinify(f"{data['name']}Event : {data['description']}")
    if not description.endswith("."):
        description += "."
    if data.get("since"):
        description += f"\n// Since obs-websocket version: {data['since'].capitalize()}."

    return f"""\
    {description}
    // {doc}#{data["heading"]["text"].lower()}
    {struct}

    // Type returns the event's update type.
    func (e {data["name"]}Event) Type() string {{ return e.UpdateType }}

    // StreamTC returns the event's stream timecode.
    func (e {data["name"]}Event) StreamTC() string {{ return e.StreamTimecode }}

    // RecTC returns the event's recording timecode.
    func (e {data["name"]}Event) RecTC() string {{ return e.RecTimecode }}
    """


def gen_requests(requests: Dict):
    """Generate all requests and responses."""
    for category, data in requests.items():
        gen_category("requests", category, data)


def gen_request(data: Dict) -> str:
    """Write Go code with type definitions and interface functions."""
    reserved = ["ID", "Type"]
    if data.get("params"):
        struct = f"""\
        type {data["name"]}Request struct {{
            {go_struct_variables(go_variables(data["params"], reserved, tag="json"))}
            _request
        }}
        """
    else:
        struct = f"type {data['name']}Request _request"

    description = newlinify(f"{data['name']}Request : {data['description']}")
    if description and not description.endswith("."):
        description += "."
    if data.get("since"):
        description += f"\n// Since obs-websocket version: {data['since'].capitalize()}."

    request = f"""\
    {description}
    // {doc}#{data["heading"]["text"].lower()}
    {struct}

    {gen_request_new(data)}

    // ID returns the request's message ID.
    func (r {data["name"]}Request) ID() string {{ return r.MessageID }}

    // Type returns the request's message type.
    func (r {data["name"]}Request) Type() string {{ return r.RequestType }}
    """

    if data.get("returns"):
        reserved = ["ID", "Stat", "Err"]
        struct = f"""\
        type {data["name"]}Response struct {{
            {go_struct_variables(go_variables(data["returns"], reserved))}
            _response `mapstructure:",squash"`
        }}
        """
    else:
        struct = f"type {data['name']}Response _response"

    description = f"// {data['name']}Response : Response for {data['name']}Request."
    if data.get("since"):
        description += f"\n// Since obs-websocket version: {data['since'].capitalize()}."

    response = f"""\
    {description}
    // {doc}#{data["heading"]["text"].lower()}
    {struct}

    // ID returns the response's message ID.
    func (r {data["name"]}Response) ID() string {{ return r.MessageID }}

    // Stat returns the response's status.
    func (r {data["name"]}Response) Stat() string {{ return r.Status }}

    // Err returns the response's error.
    func (r {data["name"]}Response) Err() string {{ return r.Error }}
    """

    return f"{request}\n\n{response}"


def gen_request_new(request: Dict):
    """Generate Go code with a New___Request function for a request type."""
    base = f"""\
    // New{request["name"]}Request returns a new {request["name"]}Request.
    func (c *Client) New{request["name"]}Request(\
    """
    variables = go_variables(request.get("params", []), [], export=False)
    if not variables:
        sig = f"{base}) {request['name']}Request {{"
        constructor_args = f'{{MessageID: c.getMessageID(), RequestType: "{request["name"]}"}}'
    else:
        args = "\n".join(
            f"{'_type' if var['name'] == 'type' else var['name']} {var['type']},"
            for var in variables
        )
        constructor_args = "{\n" + "\n".join(
            "_type," if var["name"] == "type" else f"{var['name']},"
            for var in variables
        ) + f"""
        _request{{
            MessageID: c.getMessageID(),
            RequestType: "{request["name"]}",
        }},
        }}
        """
        if len(variables) == 1:
            sig = f"{base}{args}) {request['name']}Request {{"
        else:
            sig = f"""\
            {base}
                {args}
            ) {request["name"]}Request {{\
            """
    return f"""\
    {sig}
        return {request["name"]}Request{constructor_args}
    }}\
    """


def gen_typeswitches(data: Dict):
    """Generate a Go file with a mappings from type names to structs."""
    resp_map = {}
    for category in data["requests"].values():
        for r in category:
            resp_map[r["name"]] = f"&{r['name']}Response{{}}"
    map_entries = "\n".join(f'"{k}": {v},' for k, v in resp_map.items())

    event_map = {}
    for category in data["events"].values():
        for e in category:
            event_map[e["name"]] = f"&{e['name']}Event{{}}"
    event_entries = "\n".join(f'"{k}": {v},' for k, v in event_map.items())

    switch_list = []
    for resp in resp_map:
        switch_list.append(f"""\
        case *{resp}Response:
            return *r\
        """)
    switch_entries = "\n".join(switch_list)

    with open("typeswitches.go", "w") as f:
        f.write(f"""\
        package {package}

        {disclaimer}

        var respMap = map[string]response{{
            {map_entries}
        }}

        var eventMap = map[string]Event{{
            {event_entries}
        }}

        func deref(r response) response {{
            switch r := r.(type) {{
            {switch_entries}
            default:
                return nil
            }}
        }}
        """)


def go_variables(
        variables: List[Dict],
        reserved: List[str],
        tag: str = "mapstructure",
        export: bool = True,
) -> str:
    """
    Convert a list of variable names into Go code to be put
    inside a struct definition.
    """
    vardicts, varnames = [], []
    for v in variables:
        typename, optional = optional_type(v["type"])
        varname = go_var(v["name"], export=export)
        vardicts.append({
            "name": varname,
            "type": type_map[typename.lower()],
            "tag": f'`{tag}:"{v["name"]}"`',
            "description": v["description"].replace("\n", " "),
            "optional": optional,
            "unknown": typename.lower() in unknown_types,
            "actual_type": v["type"],
            "duplicate": varname in varnames,
            "reserved": varname in reserved,
        })
        varnames.append(varname)
    return vardicts


def go_var(s: str, export: bool = True) -> str:
    """Convert a variable name in the input file to a Go variable name."""
    s = f"{(str.upper if export else str.lower)(s[0])}{s[1:]}"
    for sep in ["-", "_", ".*.", "[].", "."]:
        while sep in s:
            _len = len(sep)
            if s.endswith(sep):
                s = s[:-_len]
                continue
            i = s.find(sep)
            s = f"{s[:i]}{s[i+_len].upper()}{s[i+_len+1:]}"

    return s.replace("Id", "ID").replace("Obs", "OBS")


def go_struct_variables(variables: List[Dict]) -> str:
    """Generate Go code containing struct field definitions."""
    lines = []
    for var in variables:
        if var["description"]:
            description = var["description"]\
                          .replace("e.g. ", "e.g.")\
                          .replace(". ", "\n")\
                          .replace("e.g.", "e.g. ")
            for desc_line in description.split("\n"):
                desc_line = desc_line.strip()
                if desc_line and not desc_line.endswith("."):
                    desc_line += "."
                lines.append(f"// {desc_line}")
        lines.append(f"// Required: {'Yes' if not var['optional'] else 'No'}.")
        todos = []
        if var["unknown"]:
            todos.append(f"Unknown type ({var['actual_type']})")
        if var["duplicate"]:
            todos.append("Duplicate name")
        if var["reserved"]:
            todos.append("Reserved name")
        todos = " ".join(f"TODO: {todo}." for todo in todos)
        if todos:
            lines.append(f"// {todos}")
        lines.append(f"{var['name']} {var['type']} {var['tag']}")
    return "\n".join(lines)


def newlinify(s: str, comment: bool = True) -> str:
    """Put each sentence of a string onto its own line."""
    s = s.replace("e.g. ", "e.g.").replace(". ", "\n").replace("e.g.", "e.g. ")
    if comment:
        s = "\n".join([f"// {_s}" if not _s.startswith("//") else _s for _s in s.split("\n")])
    return s


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Missing filename argument")
        exit(1)

    if not os.path.isfile(sys.argv[1]):
        print(f"file '{sys.argv[1]}' does not exist")
        exit(1)

    with open(sys.argv[1]) as f:
        d = json.load(f)

    process_json(d)
    os.system("gofmt -w *.go")
