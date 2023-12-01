package utls

import (
	"github.com/stretchr/testify/require"
	"reflect"
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

func mockSliceExample() []bool {
	return []bool{true, false, true}
}

func mockMapExample() map[string]bool {
	return map[string]bool{
		"test":  true,
		"test2": false,
	}
}

func mockStructExample() toyStruct {
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
	sliceExample := mockSliceExample()
	structExample := mockStructExample()
	mapExample := mockMapExample()

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
			name: "map",
			item: mapExample,
		},
		{
			name: "map pointer",
			item: &mapExample,
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
			case map[string]bool:
				mapItem := tc.item.(map[string]bool)
				mapPtr := ToPtr(mapItem)
				require.Equal(t, mapItem, *mapPtr)
			case *map[string]bool:
				mapPtrItem := tc.item.(*map[string]bool)
				mapPtrPtr := ToPtr(mapPtrItem)
				require.Equal(t, mapPtrItem, *mapPtrPtr)
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

func TestToVal(t *testing.T) {
	intExample := 5
	strExample := "test"
	boolExample := true
	sliceExample := mockSliceExample()
	mapExample := mockMapExample()
	structExample := mockStructExample()
	intPtrExample := &intExample

	var intPtrPtrNilExample **int
	var intPtrNil *int

	testCases := []struct {
		name        string
		ptr         any
		expectedVal any
		ok          bool
	}{
		{
			name:        "integer pointer",
			ptr:         &intExample,
			expectedVal: intExample,
			ok:          true,
		},
		{
			name:        "string pointer",
			ptr:         &strExample,
			expectedVal: strExample,
			ok:          true,
		},
		{
			name:        "boolean pointer",
			ptr:         &boolExample,
			expectedVal: boolExample,
			ok:          true,
		},
		{
			name:        "slice pointer",
			ptr:         &sliceExample,
			expectedVal: sliceExample,
			ok:          true,
		},
		{
			name:        "map pointer",
			ptr:         &mapExample,
			expectedVal: mapExample,
			ok:          true,
		},
		{
			name:        "struct pointer",
			ptr:         &structExample,
			expectedVal: structExample,
			ok:          true,
		},
		{
			name:        "integer pointer pointer",
			ptr:         &intPtrExample,
			expectedVal: intPtrExample,
			ok:          true,
		},
		{
			name:        "integer pointer nil",
			ptr:         (*int)(nil),
			expectedVal: 0,
			ok:          false,
		},
		{
			name:        "string pointer nil",
			ptr:         (*string)(nil),
			expectedVal: "",
			ok:          false,
		},
		{
			name:        "boolean pointer nil",
			ptr:         (*bool)(nil),
			expectedVal: false,
			ok:          false,
		},
		{
			name:        "slice pointer nil",
			ptr:         (*[]bool)(nil),
			expectedVal: []bool(nil),
			ok:          false,
		},
		{
			name:        "map pointer nil",
			ptr:         (*map[string]bool)(nil),
			expectedVal: map[string]bool(nil),
			ok:          false,
		},
		{
			name:        "struct pointer nil",
			ptr:         (*toyStruct)(nil),
			expectedVal: toyStruct{},
			ok:          false,
		},
		{
			name:        "integer pointer pointer nil",
			ptr:         intPtrPtrNilExample,
			expectedVal: intPtrNil,
			ok:          false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.ptr.(type) {
			case *int:
				intPtr := tc.ptr.(*int)
				intVal, intOk := ToVal(intPtr)
				require.Equal(t, tc.ok, intOk)
				require.Equal(t, tc.expectedVal, intVal)
			case *string:
				strPtr := tc.ptr.(*string)
				strVal, strOk := ToVal(strPtr)
				require.Equal(t, tc.ok, strOk)
				require.Equal(t, tc.expectedVal, strVal)
			case *bool:
				boolPtr := tc.ptr.(*bool)
				boolVal, boolOk := ToVal(boolPtr)
				require.Equal(t, tc.ok, boolOk)
				require.Equal(t, tc.expectedVal, boolVal)
			case *[]bool:
				slicePtr := tc.ptr.(*[]bool)
				sliceVal, sliceOk := ToVal(slicePtr)
				require.Equal(t, tc.ok, sliceOk)
				require.Equal(t, tc.expectedVal, sliceVal)
			case *map[string]bool:
				mapPtr := tc.ptr.(*map[string]bool)
				mapVal, mapOk := ToVal(mapPtr)
				require.Equal(t, tc.ok, mapOk)
				require.Equal(t, tc.expectedVal, mapVal)
			case *toyStruct:
				structPtr := tc.ptr.(*toyStruct)
				structVal, structOk := ToVal(structPtr)
				require.Equal(t, tc.ok, structOk)
				require.Equal(t, tc.expectedVal, structVal)
			case **int:
				intPtrPtr := tc.ptr.(**int)
				intVal, intOk := ToVal(intPtrPtr)
				require.Equal(t, tc.ok, intOk)
				require.True(t, reflect.DeepEqual(tc.expectedVal.(*int), intVal))
				require.Equal(t, tc.expectedVal.(*int), intVal)
			default:
				t.Errorf("unexpected type: %T", tc.ptr)
			}
		})
	}
}

func TestSliceContains(t *testing.T) {
	intExample := 5
	intExample2 := 2
	strExample := "test"
	strExample2 := "test2"
	boolExample := true
	boolExample2 := false

	testCases := []struct {
		name string
		s    []any
		item any
		ok   bool
	}{
		{
			name: "integer",
			s: []any{
				intExample,
			},
			item: intExample,
			ok:   true,
		},
		{
			name: "integer pointer",
			s: []any{
				&intExample,
			},
			item: &intExample,
			ok:   true,
		},
		{
			name: "string",
			s: []any{
				strExample,
			},
			item: strExample,
			ok:   true,
		},
		{
			name: "string pointer",
			s: []any{
				&strExample,
			},
			item: &strExample,
			ok:   true,
		},
		{
			name: "boolean",
			s: []any{
				boolExample,
			},
			item: boolExample,
			ok:   true,
		},
		{
			name: "boolean pointer",
			s: []any{
				&boolExample,
			},
			item: &boolExample,
			ok:   true,
		},
		{
			name: "integer not in slice",
			s: []any{
				intExample,
			},
			item: intExample2,
			ok:   false,
		},
		{
			name: "integer pointer not in slice",
			s: []any{
				&intExample,
			},
			item: &intExample2,
			ok:   false,
		},
		{
			name: "string not in slice",
			s: []any{
				strExample,
			},
			item: strExample2,
			ok:   false,
		},
		{
			name: "string pointer not in slice",
			s: []any{
				&strExample,
			},
			item: &strExample2,
			ok:   false,
		},
		{
			name: "boolean not in slice",
			s: []any{
				boolExample,
			},
			item: boolExample2,
			ok:   false,
		},
		{
			name: "boolean pointer not in slice",
			s: []any{
				&boolExample,
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
				var intSlice []int
				for _, v := range tc.s {
					intSlice = append(intSlice, v.(int))
				}
				intOk := SliceContains(intSlice, intItem)
				require.Equal(t, tc.ok, intOk)
			case *int:
				intPtrItem := tc.item.(*int)
				var intPtrSlice []*int
				for _, v := range tc.s {
					intPtrSlice = append(intPtrSlice, v.(*int))
				}
				intPtrOk := SliceContains(intPtrSlice, intPtrItem)
				require.Equal(t, tc.ok, intPtrOk)
			case string:
				stringItem := tc.item.(string)
				var stringSlice []string
				for _, v := range tc.s {
					stringSlice = append(stringSlice, v.(string))
				}
				stringOk := SliceContains(stringSlice, stringItem)
				require.Equal(t, tc.ok, stringOk)
			case *string:
				stringPtrItem := tc.item.(*string)
				var stringPtrSlice []*string
				for _, v := range tc.s {
					stringPtrSlice = append(stringPtrSlice, v.(*string))
				}
				stringPtrOk := SliceContains(stringPtrSlice, stringPtrItem)
				require.Equal(t, tc.ok, stringPtrOk)
			case bool:
				boolItem := tc.item.(bool)
				var boolSlice []bool
				for _, v := range tc.s {
					boolSlice = append(boolSlice, v.(bool))
				}
				boolOk := SliceContains(boolSlice, boolItem)
				require.Equal(t, tc.ok, boolOk)
			case *bool:
				boolPtrItem := tc.item.(*bool)
				var boolPtrSlice []*bool
				for _, v := range tc.s {
					boolPtrSlice = append(boolPtrSlice, v.(*bool))
				}
				boolPtrOk := SliceContains(boolPtrSlice, boolPtrItem)
				require.Equal(t, tc.ok, boolPtrOk)
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

func TestSliceToMap(t *testing.T) {
	intExample := 5
	intExample2 := 2
	strExample := "test"
	strExample2 := "test2"
	boolExample := true
	boolExample2 := false

	testCases := []struct {
		name string
		s    []any
		m    map[any]bool
	}{
		{
			name: "integer",
			s: []any{
				intExample,
				intExample2,
			},
			m: map[any]bool{
				intExample:  true,
				intExample2: true,
			},
		},
		{
			name: "integer pointer",
			s: []any{
				&intExample,
				&intExample2,
			},
			m: map[any]bool{
				&intExample:  true,
				&intExample2: true,
			},
		},
		{
			name: "string",
			s: []any{
				strExample,
				strExample2,
			},
			m: map[any]bool{
				strExample:  true,
				strExample2: true,
			},
		},
		{
			name: "string pointer",
			s: []any{
				&strExample,
				&strExample2,
			},
			m: map[any]bool{
				&strExample:  true,
				&strExample2: true,
			},
		},
		{
			name: "boolean",
			s: []any{
				boolExample,
				boolExample2,
			},
			m: map[any]bool{
				boolExample:  true,
				boolExample2: true,
			},
		},
		{
			name: "boolean pointer",
			s: []any{
				&boolExample,
				&boolExample2,
			},
			m: map[any]bool{
				&boolExample:  true,
				&boolExample2: true,
			},
		},
		{
			name: "case with multiple slice values that are the same",
			s: []any{
				"value1",
				"value2",
				"value3",
				"value2",
				"value3",
				"value3",
			},
			m: map[any]bool{
				"value1": true,
				"value2": true,
				"value3": true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.s[0].(type) {
			case int:
				var intSlice []int
				for _, v := range tc.s {
					intSlice = append(intSlice, v.(int))
				}
				intMap := make(map[int]bool)
				for k, v := range tc.m {
					intMap[k.(int)] = v
				}
				m := SliceToMap(intSlice)
				for k, _ := range m {
					require.Equal(t, intMap[k], m[k])
					require.Equal(t, intMap[k], true)
					require.True(t, SliceContains(intSlice, k))
				}
				for _, v := range intSlice {
					require.Equal(t, m[v], true)
				}
			case *int:
				var intPtrSlice []*int
				for _, v := range tc.s {
					intPtrSlice = append(intPtrSlice, v.(*int))
				}
				intPtrMap := make(map[*int]bool)
				for k, v := range tc.m {
					intPtrMap[k.(*int)] = v
				}
				m := SliceToMap(intPtrSlice)
				for k, _ := range m {
					require.Equal(t, intPtrMap[k], m[k])
					require.Equal(t, intPtrMap[k], true)
					require.True(t, SliceContains(intPtrSlice, k))
				}
				for _, v := range intPtrSlice {
					require.Equal(t, m[v], true)
				}
			case string:
				var stringSlice []string
				for _, v := range tc.s {
					stringSlice = append(stringSlice, v.(string))
				}
				stringMap := make(map[string]bool)
				for k, v := range tc.m {
					stringMap[k.(string)] = v
				}
				m := SliceToMap(stringSlice)
				for k, _ := range m {
					require.Equal(t, stringMap[k], m[k])
					require.Equal(t, stringMap[k], true)
					require.True(t, SliceContains(stringSlice, k))
				}
				for _, v := range stringSlice {
					require.Equal(t, m[v], true)
				}
			case *string:
				var stringPtrSlice []*string
				for _, v := range tc.s {
					stringPtrSlice = append(stringPtrSlice, v.(*string))
				}
				stringPtrMap := make(map[*string]bool)
				for k, v := range tc.m {
					stringPtrMap[k.(*string)] = v
				}
				m := SliceToMap(stringPtrSlice)
				for k, _ := range m {
					require.Equal(t, stringPtrMap[k], m[k])
					require.Equal(t, stringPtrMap[k], true)
					require.True(t, SliceContains(stringPtrSlice, k))
				}
				for _, v := range stringPtrSlice {
					require.Equal(t, m[v], true)
				}
			case bool:
				var boolSlice []bool
				for _, v := range tc.s {
					boolSlice = append(boolSlice, v.(bool))
				}
				boolMap := make(map[bool]bool)
				for k, v := range tc.m {
					boolMap[k.(bool)] = v
				}
				m := SliceToMap(boolSlice)
				for k, _ := range m {
					require.Equal(t, boolMap[k], m[k])
					require.Equal(t, boolMap[k], true)
					require.True(t, SliceContains(boolSlice, k))
				}
				for _, v := range boolSlice {
					require.Equal(t, m[v], true)
				}
			case *bool:
				var boolPtrSlice []*bool
				for _, v := range tc.s {
					boolPtrSlice = append(boolPtrSlice, v.(*bool))
				}
				boolPtrMap := make(map[*bool]bool)
				for k, v := range tc.m {
					boolPtrMap[k.(*bool)] = v
				}
				m := SliceToMap(boolPtrSlice)
				for k, _ := range m {
					require.Equal(t, boolPtrMap[k], m[k])
					require.Equal(t, boolPtrMap[k], true)
					require.True(t, SliceContains(boolPtrSlice, k))
				}
				for _, v := range boolPtrSlice {
					require.Equal(t, m[v], true)
				}
			default:
				t.Errorf("unexpected type: %T", tc.s[0])
			}
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name string
		a    any
		b    any
		min  any
	}{
		{
			name: "integer a > b",
			a:    5,
			b:    2,
			min:  2,
		},
		{
			name: "integer a < b",
			a:    2,
			b:    5,
			min:  2,
		},
		{
			name: "integer a == b",
			a:    5,
			b:    5,
			min:  5,
		},
		{
			name: "float a > b",
			a:    5.0,
			b:    2.0,
			min:  2.0,
		},
		{
			name: "float a < b",
			a:    2.0,
			b:    5.0,
			min:  2.0,
		},
		{
			name: "float a == b",
			a:    5.0,
			b:    5.0,
			min:  5.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.a.(type) {
			case int:
				intA := tc.a.(int)
				intB := tc.b.(int)
				intMin := Min(intA, intB)
				require.Equal(t, tc.min, intMin)
			case float64:
				floatA := tc.a.(float64)
				floatB := tc.b.(float64)
				floatMin := Min(floatA, floatB)
				require.Equal(t, tc.min, floatMin)
			default:
				t.Errorf("unexpected type: %T", tc.a)
			}
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name string
		a    any
		b    any
		max  any
	}{
		{
			name: "integer a > b",
			a:    5,
			b:    2,
			max:  5,
		},
		{
			name: "integer a < b",
			a:    2,
			b:    5,
			max:  5,
		},
		{
			name: "integer a == b",
			a:    5,
			b:    5,
			max:  5,
		},
		{
			name: "float a > b",
			a:    5.0,
			b:    2.0,
			max:  5.0,
		},
		{
			name: "float a < b",
			a:    2.0,
			b:    5.0,
			max:  5.0,
		},
		{
			name: "float a == b",
			a:    5.0,
			b:    5.0,
			max:  5.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.a.(type) {
			case int:
				intA := tc.a.(int)
				intB := tc.b.(int)
				intMax := Max(intA, intB)
				require.Equal(t, tc.max, intMax)
			case float64:
				floatA := tc.a.(float64)
				floatB := tc.b.(float64)
				floatMax := Max(floatA, floatB)
				require.Equal(t, tc.max, floatMax)
			default:
				t.Errorf("unexpected type: %T", tc.a)
			}
		})
	}
}
