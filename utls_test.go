package utls

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToPtr(t *testing.T) {
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

	intExample := 5
	strExample := "test"
	boolExample := true
	sliceExample := []bool{true, false, true}
	structExample := toyStruct{
		i:      5,
		iPtr:   ToPtr(5),
		s:      "test",
		sPtr:   ToPtr("test"),
		b:      true,
		bPtr:   ToPtr(true),
		slc:    []bool{true, false, true},
		slcPtr: ToPtr([]bool{true, false, true}),
	}

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
