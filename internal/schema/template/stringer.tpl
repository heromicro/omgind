{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "stringer" }}

{{/* Add the base header for the generated file */}}
{{ $pkg := base $.Config.Package }}
{{ template "header" $ }}

{{/* Loop over all nodes and implement the "GoStringer" interface */}}
{{ range $n := $.Nodes }}
    {{ $receiver := $n.Receiver }}
    func ({{ $receiver }} *{{ $n.Name }}) GoString() string {
        if {{ $receiver }} == nil {
            return fmt.Sprintf("{{ $n.Name }}(nil)")
        }
        return {{ $receiver }}.String()
    }
{{ end }}

{{ end }}
