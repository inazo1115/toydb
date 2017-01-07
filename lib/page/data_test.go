package page

import (
	//"fmt"
	"testing"

	"github.com/inazo1115/toydb/lib/util"
)

func TestAddRecord(t *testing.T) {

	expected := "0123456789"

	dataPage := NewDataPage(-1, -1, -1, int64(len(expected)))
	if err := dataPage.AddRecord([]byte(expected)); err != nil {
		t.Errorf("AddRecord failed.")
	}

	actual := string(dataPage.Data()[:10])
	util.Assert(t, actual, expected)
}

func TestReadRecord(t *testing.T) {

	expected0 := "0123456789"
	expected1 := "helloworld"

	dataPage := NewDataPage(-1, -1, -1, int64(len(expected0)))
	if err := dataPage.AddRecord([]byte(expected0)); err != nil {
		t.Errorf("AddRecord failed.")
	}
	if err := dataPage.AddRecord([]byte(expected1)); err != nil {
		t.Errorf("AddRecord failed.")
	}

	actual0 := string(dataPage.ReadRecord(0))
	actual1 := string(dataPage.ReadRecord(1))

	util.Assert(t, actual0, expected0)
	util.Assert(t, actual1, expected1)
}

func TestSerde(t *testing.T) {

	dataPage0 := NewDataPage(3, 2, 4, 10)
	if err := dataPage0.AddRecord([]byte("0123456789")); err != nil {
		t.Errorf("AddRecord failed.")
	}

	ser, err := dataPage0.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary failed.")
	}

	dataPage1 := &DataPage{}
	err = dataPage1.UnmarshalBinary(ser)
	if err != nil {
		t.Errorf("UnmarshalBinary failed.")
	}

	util.Assert(t, dataPage1.Pid(), dataPage0.Pid())
	util.Assert(t, dataPage1.Previous(), dataPage0.Previous())
	util.Assert(t, dataPage1.Next(), dataPage0.Next())
	util.Assert(t, dataPage1.NumRecords(), dataPage0.NumRecords())
	util.Assert(t, string(dataPage1.ReadRecord(0)), string(dataPage0.ReadRecord(0)))
	util.Assert(t, string(dataPage1.ReadRecord(1)), string(dataPage0.ReadRecord(1)))
}
