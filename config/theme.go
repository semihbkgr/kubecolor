package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kubecolor/kubecolor/internal/stringutil"
	"github.com/spf13/viper"
)

// NewBaseTheme returns the base color schema depending on the dark/light setting
func NewBaseTheme(preset Preset) *Theme {
	switch preset {
	case PresetDark:
		return &Theme{
			Default: MustParseColor("green"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("hicyan / cyan"),
				Info:      MustParseColor("white"),
				Primary:   MustParseColor("magenta"),
				Secondary: MustParseColor("cyan"),
				Success:   MustParseColor("green"),
				Warning:   MustParseColor("yellow"),
				Danger:    MustParseColor("red"),
				Muted:     MustParseColor("gray:italic"),
			},
			Table: ThemeTable{
				Header: MustParseColor("bold"),
			},
			Data: ThemeData{
				String: MustParseColor("hiyellow"),
			},
		}

	case PresetLight:
		return &Theme{
			Default: MustParseColor("green"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("hiblue / blue"),
				Info:      MustParseColor("black"),
				Primary:   MustParseColor("magenta"),
				Secondary: MustParseColor("blue"),
				Success:   MustParseColor("green"),
				Warning:   MustParseColor("yellow"),
				Danger:    MustParseColor("red"),
				Muted:     MustParseColor("gray:italic"),
			},
			Table: ThemeTable{
				Header: MustParseColor("bold"),
			},
			Data: ThemeData{
				String: MustParseColor("yellow"),
			},
		}

	// Special Preset for Protanopias
	case PresetProtDark:
		return &Theme{
			Default: MustParseColor("white"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("#feb927 / #fe6e1a"),
				Info:      MustParseColor("white"),
				Primary:   MustParseColor("#4860e6"),             // magenta
				Secondary: MustParseColor("#2aabee"),             // cyan
				Success:   MustParseColor("#6afd6a:bold"),        // bold green
				Warning:   MustParseColor("#feb927:italic"),      // yellow
				Danger:    MustParseColor("fg=white:bg=#c2270a"), // red background
				Muted:     MustParseColor("#2ee5ae:italic"),      // white-ish
			},
			Data: ThemeData{
				String: MustParseColor("#2aabee"),
			},
			Table: ThemeTable{
				Header:  MustParseColor("white:bold"),
				Columns: MustParseColorSlice("#2aabee / #6afd6a:bold / #4860e6 / white / #feb927"),
			},
		}

	case PresetProtLight:
		return &Theme{
			Default: MustParseColor("black"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("#feb927 / #fe6e1a"),
				Info:      MustParseColor("black"),
				Primary:   MustParseColor("#4860e6"),
				Secondary: MustParseColor("#2aabee"),
				Success:   MustParseColor("#6afd6a:bold"),
				Warning:   MustParseColor("#feb927:italic"),
				Danger:    MustParseColor("fg=black:bg=#c2270a"),
				Muted:     MustParseColor("#2ee5ae:italic"),
			},
			Data: ThemeData{
				String: MustParseColor("#2aabee"),
			},
			Table: ThemeTable{
				Header:  MustParseColor("black:bold"),
				Columns: MustParseColorSlice("#2aabee / #6afd6a:bold / #4860e6 / black / #feb927"),
			},
		}

	// Special Preset for Deuteranopia
	case PresetDeutDark:
		return &Theme{
			Default: MustParseColor("white"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("#feb927 / #fe6e1a"),
				Info:      MustParseColor("white"),
				Primary:   MustParseColor("#4860e6"),
				Secondary: MustParseColor("#2aabee"),
				Success:   MustParseColor("#6afd6a:bold"),
				Warning:   MustParseColor("#feb927:italic"),
				Danger:    MustParseColor("fg=white:bg=#c2270a"),
				Muted:     MustParseColor("#2ee5ae"),
			},
			Data: ThemeData{
				String: MustParseColor("#2aabee"),
			},
			Table: ThemeTable{
				Header:  MustParseColor("white:bold"),
				Columns: MustParseColorSlice("#2aabee / #6afd6a:bold / #4860e6 / white / #feb927"),
			},
		}

	case PresetDeutLight:
		return &Theme{
			Default: MustParseColor("black"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("#feb927 / #fe6e1a"),
				Info:      MustParseColor("black"),
				Primary:   MustParseColor("#4860e6"),
				Secondary: MustParseColor("#2aabee"),
				Success:   MustParseColor("#6afd6a:bold"),
				Warning:   MustParseColor("#feb927:italic"),
				Danger:    MustParseColor("fg=black:bg=#c2270a"),
				Muted:     MustParseColor("#2ee5ae:italic"),
			},
			Data: ThemeData{
				String: MustParseColor("#2aabee"),
			},
			Table: ThemeTable{
				Header:  MustParseColor("black:bold"),
				Columns: MustParseColorSlice("#2aabee / #6afd6a:bold / #4860e6 / black / #feb927"),
			},
		}

	// Special Preset for Tritanopia
	case PresetTritDark:
		return &Theme{
			Default: MustParseColor("white"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("#feb927 / #fe6e1a"),
				Info:      MustParseColor("white"),
				Primary:   MustParseColor("#4860e6"),
				Secondary: MustParseColor("#2aabee"),
				Success:   MustParseColor("#6afd6a:bold"),
				Warning:   MustParseColor("#feb927:italic"),
				Danger:    MustParseColor("fg=white:bg=#c2270a"),
				Muted:     MustParseColor("#2ee5ae"),
			},
			Data: ThemeData{
				String: MustParseColor("#2aabee"),
			},
			Table: ThemeTable{
				Header:  MustParseColor("white:bold"),
				Columns: MustParseColorSlice("#2aabee / #6afd6a:bold / #4860e6 / white / #feb927"),
			},
		}

	case PresetTritLight:
		return &Theme{
			Default: MustParseColor("black"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("#feb927 / #fe6e1a"),
				Info:      MustParseColor("black"),
				Primary:   MustParseColor("#4860e6"),
				Secondary: MustParseColor("#2aabee"),
				Success:   MustParseColor("#6afd6a:bold"),
				Warning:   MustParseColor("#feb927:italic"),
				Danger:    MustParseColor("fg=black:bg=#c2270a"),
				Muted:     MustParseColor("#2ee5ae:italic"),
			},
			Data: ThemeData{
				String: MustParseColor("#2aabee"),
			},
			Table: ThemeTable{
				Header:  MustParseColor("black:bold"),
				Columns: MustParseColorSlice("#2aabee / #6afd6a:bold / #4860e6 / black / #feb927"),
			},
		}

	// Pre-v0.3.0
	case PresetPre030Dark:
		return &Theme{
			Default: MustParseColor("green"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("yellow / white"),
				Info:      MustParseColor("white"),
				Primary:   MustParseColor("magenta"),
				Secondary: MustParseColor("cyan"),
				Success:   MustParseColor("green"),
				Warning:   MustParseColor("yellow"),
				Danger:    MustParseColor("red"),
				Muted:     MustParseColor("yellow"),
			},
			Options: ThemeOptions{
				Flag: MustParseColor("yellow"),
			},
		}

	case PresetPre030Light:
		return &Theme{
			Default: MustParseColor("green"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("yellow / black"),
				Info:      MustParseColor("black"),
				Primary:   MustParseColor("magenta"),
				Secondary: MustParseColor("blue"),
				Success:   MustParseColor("green"),
				Warning:   MustParseColor("yellow"),
				Danger:    MustParseColor("red"),
				Muted:     MustParseColor("yellow"),
			},
			Options: ThemeOptions{
				Flag: MustParseColor("yellow"),
			},
		}

	// Pre-v0.0.21
	case PresetPre0021Dark:
		return &Theme{
			Default: MustParseColor("green"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("yellow / white"),
				Info:      MustParseColor("white"),
				Primary:   MustParseColor("magenta"),
				Secondary: MustParseColor("cyan"),
				Success:   MustParseColor("green"),
				Warning:   MustParseColor("yellow"),
				Danger:    MustParseColor("red"),
				Muted:     MustParseColor("yellow"),
			},
			Data: ThemeData{
				String: MustParseColor("cyan"),
			},
			Table: ThemeTable{
				Columns: MustParseColorSlice("cyan / green / magenta / white / yellow"),
			},
			Status: ThemeStatus{
				Success: MustParseColor("none"),
			},
			Options: ThemeOptions{
				Flag: MustParseColor("yellow"),
			},
		}

	case PresetPre0021Light:
		return &Theme{
			Default: MustParseColor("green"),
			Base: ThemeBase{
				Key:       MustParseColorSlice("yellow / black"),
				Info:      MustParseColor("white"),
				Primary:   MustParseColor("magenta"),
				Secondary: MustParseColor("cyan"),
				Success:   MustParseColor("green"),
				Warning:   MustParseColor("yellow"),
				Danger:    MustParseColor("red"),
				Muted:     MustParseColor("yellow"),
			},
			Data: ThemeData{
				String: MustParseColor("blue"),
			},
			Table: ThemeTable{
				Columns: MustParseColorSlice("cyan / green / magenta / black / yellow / blue"),
			},
			Status: ThemeStatus{
				Success: MustParseColor("none"),
			},
			Options: ThemeOptions{
				Flag: MustParseColor("yellow"),
			},
		}

	default:
		// Empty theme
		return &Theme{}
	}
}

// Theme is the root theme config.
type Theme struct {
	// Base colors must be first so they're applied first
	Base ThemeBase // base colors for themes

	Default Color // default when no specific mapping is found for the command

	Shell  ThemeShell  // colors for representing shells (e.g bash, zsh, etc)
	Data   ThemeData   // colors for representing data
	Status ThemeStatus // generic status coloring (e.g "Ready", "Terminating")
	Table  ThemeTable  // used in table output, e.g "kubectl get" and parts of "kubectl describe"
	Stderr ThemeStderr // used in kubectl's stderr output

	Apply    ThemeApply    // used in "kubectl apply"
	Create   ThemeCreate   // used in "kubectl create"
	Delete   ThemeDelete   // used in "kubectl delete"
	Describe ThemeDescribe // used in "kubectl describe"
	Drain    ThemeDrain    // used in "kubectl drain"
	Explain  ThemeExplain  // used in "kubectl explain"
	Expose   ThemeExpose   // used in "kubectl expose"
	Options  ThemeOptions  // used in "kubectl options"
	Patch    ThemePatch    // used in "kubectl patch"
	Rollout  ThemeRollout  // used in "kubectl rollout"
	Scale    ThemeScale    // used in "kubectl scale"
	Uncordon ThemeUncordon // used in "kubectl uncordon"
	Version  ThemeVersion  // used in "kubectl version"
	Help     ThemeHelp     // used in "kubectl --help"
	Logs     ThemeLogs     // used in "kubectl logs"
}

func (t *Theme) ComputeCache() {
	themeVal := reflect.ValueOf(t).Elem()
	walkFields(themeVal, "theme", visitorComputeCache)
}

// ThemeBase contains base colors that other theme fields can default to,
// just to make overriding themes easier.
//
// These fields should never be referenced in the printers.
// Instead, they should use the more specific fields, such as [ThemeApply.Created]
type ThemeBase struct {
	Info      Color // general color for when things are informational
	Primary   Color // general color for when things are focus
	Secondary Color // general color for when things are secondary focus
	Success   Color // general color for when things are good
	Warning   Color // general color for when things are wrong
	Danger    Color // general color for when things are bad
	Muted     Color // general color for when things are less relevant

	Key ColorSlice `defaultFromMany:"theme.base.secondary"` // general color for keys
}

// ThemeShell holds colors for when representing shell commands (bash, zsh, etc)
type ThemeShell struct {
	Comment Color `defaultFrom:"theme.base.muted"`     // used on comments, e.g `# this is a comment`
	Command Color `defaultFrom:"theme.base.success"`   // used on commands, e.g `kubectl` or `echo`
	Arg     Color `defaultFrom:"theme.base.info"`      // used on arguments, e.g `get pods` in `kubectl get pods`
	Flag    Color `defaultFrom:"theme.base.secondary"` // used on flags, e.g `--watch` in `kubectl get pods --watch`
}

// ThemeData holds colors for when representing parsed data.
// Such as in YAML, JSON, and even some "kubectl describe" values
type ThemeData struct {
	Key    ColorSlice `defaultFrom:"theme.base.key"`     // used for the key
	String Color      `defaultFrom:"theme.base.info"`    // used when value is a string
	True   Color      `defaultFrom:"theme.base.success"` // used when value is true
	False  Color      `defaultFrom:"theme.base.danger"`  // used when value is false
	Number Color      `defaultFrom:"theme.base.primary"` // used when the value is a number
	Null   Color      `defaultFrom:"theme.base.muted"`   // used when the value is null, nil, or none

	Quantity      Color `defaultFrom:"theme.data.number"`  // used when the value is a quantity, e.g "100m" or "5Gi"
	Duration      Color ``                                 // used when the value is a duration, e.g "12m" or "1d12h"
	DurationFresh Color `defaultFrom:"theme.base.success"` // color used when the time value is under a certain delay

	Ratio ThemeDataRatio
}

type ThemeDataRatio struct {
	Zero    Color `defaultFrom:"theme.base.muted"`   // used for "0/0"
	Equal   Color ``                                 // used for "n/n", e.g "1/1"
	Unequal Color `defaultFrom:"theme.base.warning"` // used for "n/m", e.g "0/1"
}

// ThemeStatus holds colors for status texts, used in for example
// the "kubectl get" status column
type ThemeStatus struct {
	Success Color `defaultFrom:"theme.base.success"` // used in status keywords, e.g "Running", "Ready"
	Warning Color `defaultFrom:"theme.base.warning"` // used in status keywords, e.g "Terminating"
	Error   Color `defaultFrom:"theme.base.danger"`  // used in status keywords, e.g "Failed", "Unhealthy"
}

// ThemeTable holds colors for table output
type ThemeTable struct {
	Header  Color      `defaultFrom:"theme.base.info"`                          // used on table headers
	Columns ColorSlice `defaultFromMany:"theme.base.info,theme.base.secondary"` // used on table columns when no other coloring applies such as status or duration coloring. The multiple colors are cycled based on column ID, from left to right.
}

// ThemeStderr holds generic colors for kubectl's stderr output.
type ThemeStderr struct {
	Error Color `defaultFrom:"theme.base.danger"` // e.g when text contains "error"

	NoneFound          Color `defaultFrom:"theme.data.null"`   // used on table output like "No resources found"
	NoneFoundNamespace Color `defaultFrom:"theme.data.string"` // used on the namespace name of "No resources found in my-ns namespace"

	// default when no specific mapping is found for the output line
	//
	// Deprecated: This field is no longer used (since v0.4.0),
	// as the stderr logs now uses the "kubectl logs" behavior as a fallback/default coloring.
	Default Color `jsonschema_extras:"deprecated=true"` // *deprecated: this field is no longer used (since v0.4.0)*
}

// ThemeApply holds colors for the "kubectl apply" output.
type ThemeDescribe struct {
	Key ColorSlice `defaultFrom:"theme.base.key"` // used on keys. The multiple colors are cycled based on indentation.
}

// ThemeApply holds colors for the "kubectl apply" output.
type ThemeApply struct {
	Created    Color `defaultFrom:"theme.base.success"` // used on "deployment.apps/foo created"
	Configured Color `defaultFrom:"theme.base.warning"` // used on "deployment.apps/bar configured"
	Unchanged  Color `defaultFrom:"theme.base.primary"` // used on "deployment.apps/quux unchanged"

	DryRun   Color `defaultFrom:"theme.base.secondary"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.success"`   // used when outputs unknown format
}

// ThemeDelete holds colors for the "kubectl delete" output.
type ThemeDelete struct {
	Deleted Color `defaultFrom:"theme.base.danger"` // used on "deployment.apps "nginx" deleted"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.danger"`  // used when outputs unknown format
}

// ThemeCreate holds colors for the "kubectl create" output.
type ThemeCreate struct {
	Created Color `defaultFrom:"theme.base.success"` // used on "deployment.apps/foo created"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.success"` // used when outputs unknown format
}

// ThemeExpose holds colors for the "kubectl expose" output.
type ThemeExpose struct {
	Exposed Color `defaultFrom:"theme.base.primary"` // used on "deployment.apps/foo created"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.primary"` // used when outputs unknown format
}

// ThemeScale holds colors for the "kubectl scale" output.
type ThemeScale struct {
	Scaled Color `defaultFrom:"theme.base.warning"` // used on "deployment.apps/foo scaled"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.warning"` // used when outputs unknown format
}

// ThemeRollout holds colors for the "kubectl rollout" output.
type ThemeRollout struct {
	RolledBack Color `defaultFrom:"theme.base.warning"`   // used on "deployment.apps/foo rolled back"
	Paused     Color `defaultFrom:"theme.base.primary"`   // used on "deployment.apps/foo paused"
	Resumed    Color `defaultFrom:"theme.base.secondary"` // used on "deployment.apps/foo resumed"
	Restarted  Color `defaultFrom:"theme.base.warning"`   // used on "deployment.apps/foo restarted"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.warning"` // used when outputs unknown format
}

// ThemePatch holds colors for the "kubectl patch" output.
type ThemePatch struct {
	Patched Color `defaultFrom:"theme.base.warning"` // used on "deployment.apps/foo patched"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.warning"` // used when outputs unknown format
}

// ThemeUncordon holds colors for the "kubectl uncordon" output.
type ThemeUncordon struct {
	Uncordoned Color `defaultFrom:"theme.base.secondary"` // used on "node/my-worker-node-01 uncordoned"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.warning"` // used when outputs unknown format
}

// ThemeDrain holds colors for the "kubectl drain" output.
type ThemeDrain struct {
	Cordoned    Color `defaultFrom:"theme.base.primary"` // used on "node/my-worker-node-01 cordoned"
	EvictingPod Color `defaultFrom:"theme.base.muted"`   // used on "evicting pod my-namespace/my-pod"
	Evicted     Color `defaultFrom:"theme.base.warning"` // used on "pod/my-pod evicted"
	Drained     Color `defaultFrom:"theme.base.success"` // used on "node/my-worker-node-01 drained"

	DryRun   Color `defaultFrom:"theme.apply.dryrun"` // used on "(dry run)" and "(server dry run)"
	Fallback Color `defaultFrom:"theme.base.warning"` // used when outputs unknown format
}

// ThemeExplain holds colors for the "kubectl explain" output.
type ThemeExplain struct {
	Key      ColorSlice `defaultFrom:"theme.base.key"`    // used on keys. The multiple colors are cycled based on indentation.
	Required Color      `defaultFrom:"theme.base.danger"` // used on the trailing "-required-" string
}

// ThemeOptions holds colors for the "kubectl options" output.
type ThemeOptions struct {
	Flag Color `defaultFrom:"theme.base.secondary"` // e.g "--kubeconfig"
}

// ThemeVersion holds colors for the "kubectl version" output.
type ThemeVersion struct {
	Key ColorSlice `defaultFrom:"theme.base.key"` // used on the key
}

// ThemeHelp holds colors for the "kubectl --help" output.
type ThemeHelp struct {
	Header   Color `defaultFrom:"theme.table.header"`   // e.g "Examples:" or "Options:"
	Flag     Color `defaultFrom:"theme.base.secondary"` // e.g "--kubeconfig"
	FlagDesc Color `defaultFrom:"theme.base.info"`      // Flag descripion under "Options:" heading
	Url      Color `defaultFrom:"theme.base.secondary"` // e.g `[https://example.com]`
	Text     Color `defaultFrom:"theme.base.info"`      // Fallback text color
}

// ThemeLogs holds colors for the "kubectl logs" output.
type ThemeLogs struct {
	Key          ColorSlice `defaultFrom:"theme.data.key"`
	QuotedString Color      `defaultFrom:"theme.data.string"` // Used on quoted strings that are not part of a `key="value"`
	Date         Color      `defaultFrom:"theme.base.muted"`
	SourceRef    Color      `defaultFrom:"theme.base.muted"`
	GUID         Color      `defaultFrom:"theme.base.muted"`

	Severity ThemeLogsSeverity
}

// ThemeLogsSeverity holds colors for "log level severity" found in "kubectl logs" output
type ThemeLogsSeverity struct {
	Trace Color `defaultFrom:"theme.base.muted"`
	Debug Color `defaultFrom:"theme.base.muted"`
	Info  Color `defaultFrom:"theme.base.success"`
	Warn  Color `defaultFrom:"theme.base.warning"`
	Error Color `defaultFrom:"theme.base.danger"`
	Fatal Color `defaultFrom:"theme.base.danger"`
	Panic Color `defaultFrom:"theme.base.danger"`
}

func applyViperDefaults(theme *Theme, v *viper.Viper) {
	themeVal := reflect.ValueOf(theme).Elem()
	walkFields(themeVal, "theme", themeViperVisitor{viper: v}.visitorApplyDefaults)
}

type themeViperVisitor struct {
	viper *viper.Viper
}

func walkFields(val reflect.Value, viperKey string, visitor func(viperKey string, value reflect.Value, tags reflect.StructTag)) {
	typ := val.Type()
	for i := range val.NumField() {
		fieldTyp := typ.Field(i)
		if fieldTyp.Anonymous || !fieldTyp.IsExported() {
			continue
		}
		fieldVal := val.Field(i)
		// e.g "theme" + field "Default" => "theme.default"
		newViperKey := fmt.Sprintf("%s.%s", viperKey, strings.ToLower(fieldTyp.Name))
		// Only dig deeper if its a theme struct, e.g ThemeApply
		if strings.HasPrefix(fieldTyp.Type.Name(), "Theme") {
			walkFields(fieldVal, newViperKey, visitor)
			continue
		}
		visitor(newViperKey, fieldVal, fieldTyp.Tag)
	}
}

func visitorComputeCache(viperKey string, value reflect.Value, _ reflect.StructTag) {
	switch value := value.Addr().Interface().(type) {
	case *Color:
		value.ComputeCache()
	case *ColorSlice:
		value.ComputeCache()
	default:
		panic(fmt.Errorf("%s: unsupported field type: %T", viperKey, value))
	}
}

func (t themeViperVisitor) visitorApplyDefaults(viperKey string, value reflect.Value, tags reflect.StructTag) {
	switch value := value.Interface().(type) {
	case Color:
		if _, ok := tags.Lookup("defaultFromMany"); ok {
			panic(fmt.Errorf("%s: cannot use defaultFromMany tag on a Color field", viperKey))
		}
		if defaultFrom, ok := tags.Lookup("defaultFrom"); ok {
			t.setColorOrKey(viperKey, value, defaultFrom)
		} else {
			t.setColor(viperKey, value)
		}
	case ColorSlice:
		if defaultFrom, ok := tags.Lookup("defaultFrom"); ok {
			t.setColorSliceOrKey(viperKey, value, defaultFrom)
		} else if defaultFromMany, ok := tags.Lookup("defaultFromMany"); ok {
			split := stringutil.SplitAndTrimSpace(defaultFromMany, ",")
			t.setColorSliceOrManyKeys(viperKey, value, split)
		} else {
			t.setColorSlice(viperKey, value)
		}
	default:
		panic(fmt.Errorf("%s: unsupported field type: %T", viperKey, value))
	}
}

func (t themeViperVisitor) setColorOrKey(key string, value Color, otherKey string) {
	if t.setColor(key, value) {
		return
	}
	t.viper.SetDefault(key, t.viper.Get(otherKey))
}

func (t themeViperVisitor) setColor(key string, value Color) bool {
	if value.IsZero() {
		t.viper.SetDefault(key, Color{})
		return false
	}
	t.viper.SetDefault(key, value)
	return true
}

func (t themeViperVisitor) setColorSliceOrKey(key string, value ColorSlice, otherKey string) {
	if t.setColorSlice(key, value) {
		return
	}
	t.viper.SetDefault(key, t.viper.Get(otherKey))
}

func (t themeViperVisitor) setColorSliceOrManyKeys(key string, value ColorSlice, otherKeys []string) {
	if t.setColorSlice(key, value) {
		return
	}
	values := make(ColorSlice, 0, len(otherKeys))
	for _, k := range otherKeys {
		val := t.viper.Get(k)
		col, ok := val.(Color)
		if !ok {
			col = MustParseColor(fmt.Sprint(val))
		}
		if val != nil {
			values = append(values, col)
		}
	}
	if len(values) > 0 {
		t.viper.SetDefault(key, values)
	}
}

func (t themeViperVisitor) setColorSlice(key string, value ColorSlice) bool {
	if len(value) == 0 {
		t.viper.SetDefault(key, ColorSlice{})
		return false
	}
	t.viper.SetDefault(key, value)
	return true
}
