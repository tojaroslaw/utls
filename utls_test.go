package utls

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TEST FIXTURES

type toyStruct struct {
	i      int
	iPtr   *int
	s      string
	sPtr   *string
	b      bool
	bPtr   *bool
	slc    []bool
	slcPtr *[]bool
}

func sliceExample() []bool {
	return []bool{true, false, true}
}

func structExample() toyStruct {
	return toyStruct{
		i:      5,
		iPtr:   ToPtr(5),
		s:      "test",
		sPtr:   ToPtr("test"),
		b:      true,
		bPtr:   ToPtr(true),
		slc:    []bool{true, false, true},
		slcPtr: ToPtr([]bool{true, false, true}),
	}
}

// TESTS

func TestToPtr(t *testing.T) {
	intExample := 5
	strExample := "test"
	boolExample := true
	sliceExample := sliceExample()
	structExample := structExample()

	testCases := []struct {
		name string
		item any
	}{
		{
			name: "integer",
			item: intExample,
		},
		{
			name: "integer pointer",
			item: &intExample,
		},
		{
			name: "string",
			item: strExample,
		},
		{
			name: "string pointer",
			item: &strExample,
		},
		{
			name: "boolean",
			item: boolExample,
		},
		{
			name: "boolean pointer",
			item: &boolExample,
		},
		{
			name: "slice",
			item: sliceExample,
		},
		{
			name: "slice pointer",
			item: &sliceExample,
		},
		{
			name: "struct",
			item: structExample,
		},
		{
			name: "struct pointer",
			item: &structExample,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ptr := ToPtr(tc.item)

			require.Equal(t, tc.item, *ptr)

			switch tc.item.(type) {
			case int:
				intItem := tc.item.(int)
				intPtr := ToPtr(intItem)
				require.Equal(t, intItem, *intPtr)
			case *int:
				intPtrItem := tc.item.(*int)
				intPtrPtr := ToPtr(intPtrItem)
				require.Equal(t, intPtrItem, *intPtrPtr)
			case string:
				stringItem := tc.item.(string)
				stringPtr := ToPtr(stringItem)
				require.Equal(t, stringItem, *stringPtr)
			case *string:
				stringPtrItem := tc.item.(*string)
				stringPtrPtr := ToPtr(stringPtrItem)
				require.Equal(t, stringPtrItem, *stringPtrPtr)
			case bool:
				boolItem := tc.item.(bool)
				boolPtr := ToPtr(boolItem)
				require.Equal(t, boolItem, *boolPtr)
			case *bool:
				boolPtrItem := tc.item.(*bool)
				boolPtrPtr := ToPtr(boolPtrItem)
				require.Equal(t, boolPtrItem, *boolPtrPtr)
			case []bool:
				sliceItem := tc.item.([]bool)
				slicePtr := ToPtr(sliceItem)
				require.Equal(t, sliceItem, *slicePtr)
			case *[]bool:
				slicePtrItem := tc.item.(*[]bool)
				slicePtrPtr := ToPtr(slicePtrItem)
				require.Equal(t, slicePtrItem, *slicePtrPtr)
			case toyStruct:
				structItem := tc.item.(toyStruct)
				structPtr := ToPtr(structItem)
				require.Equal(t, structItem, *structPtr)
			case *toyStruct:
				structPtrItem := tc.item.(*toyStruct)
				structPtrPtr := ToPtr(structPtrItem)
				require.Equal(t, structPtrItem, *structPtrPtr)
			default:
				t.Errorf("unexpected type: %T", tc.item)
			}
		})
	}
}

func TestMapContains(t *testing.T) {
	intExample := 5
	intExample2 := 2
	strExample := "test"
	strExample2 := "test2"
	boolExample := true
	boolExample2 := false

	testCases := []struct {
		name string
		m    map[any]any
		item any
		ok   bool
	}{
		{
			name: "integer",
			m: map[any]any{
				intExample: "test",
			},
			item: intExample,
			ok:   true,
		},
		{
			name: "integer pointer",
			m: map[any]any{
				&intExample: "test",
			},
			item: &intExample,
			ok:   true,
		},
		{
			name: "string",
			m: map[any]any{
				strExample: 5,
			},
			item: strExample,
			ok:   true,
		},
		{
			name: "string pointer",
			m: map[any]any{
				&strExample: 5,
			},
			item: &strExample,
			ok:   true,
		},
		{
			name: "boolean",
			m: map[any]any{
				boolExample: 5,
			},
			item: boolExample,
			ok:   true,
		},
		{
			name: "boolean pointer",
			m: map[any]any{
				&boolExample: 5,
			},
			item: &boolExample,
			ok:   true,
		},
		{
			name: "integer not in map",
			m: map[any]any{
				intExample: "test",
			},
			item: intExample2,
			ok:   false,
		},
		{
			name: "integer pointer not in map",
			m: map[any]any{
				&intExample: "test",
			},
			item: &intExample2,
			ok:   false,
		},
		{
			name: "string not in map",
			m: map[any]any{
				strExample: 5,
			},
			item: strExample2,
			ok:   false,
		},
		{
			name: "string pointer not in map",
			m: map[any]any{
				&strExample: 5,
			},
			item: &strExample2,
			ok:   false,
		},
		{
			name: "boolean not in map",
			m: map[any]any{
				boolExample: 5,
			},
			item: boolExample2,
			ok:   false,
		},
		{
			name: "boolean pointer not in map",
			m: map[any]any{
				&boolExample: 5,
			},
			item: &boolExample2,
			ok:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.item.(type) {
			case int:
				intItem := tc.item.(int)
				m := map[int]any{}
				for k, v := range tc.m {
					m[k.(int)] = v
				}
				intOk := MapContains(m, intItem)
				require.Equal(t, tc.ok, intOk)
			case *int:
				intPtrItem := tc.item.(*int)
				m := map[*int]any{}
				for k, v := range tc.m {
					m[k.(*int)] = v
				}
				intPtrOk := MapContains(m, intPtrItem)
				require.Equal(t, tc.ok, intPtrOk)
			case string:
				stringItem := tc.item.(string)
				m := map[string]any{}
				for k, v := range tc.m {
					m[k.(string)] = v
				}
				stringOk := MapContains(m, stringItem)
				require.Equal(t, tc.ok, stringOk)
			case *string:
				stringPtrItem := tc.item.(*string)
				m := map[*string]any{}
				for k, v := range tc.m {
					m[k.(*string)] = v
				}
				stringPtrOk := MapContains(m, stringPtrItem)
				require.Equal(t, tc.ok, stringPtrOk)
			case bool:
				boolItem := tc.item.(bool)
				m := map[bool]any{}
				for k, v := range tc.m {
					m[k.(bool)] = v
				}
				boolOk := MapContains(m, boolItem)
				require.Equal(t, tc.ok, boolOk)
			case *bool:
				boolPtrItem := tc.item.(*bool)
				m := map[*bool]any{}
				for k, v := range tc.m {
					m[k.(*bool)] = v
				}
				boolPtrOk := MapContains(m, boolPtrItem)
				require.Equal(t, tc.ok, boolPtrOk)
			default:
				t.Errorf("unexpected type: %T", tc.item)
			}
		})
	}
}
