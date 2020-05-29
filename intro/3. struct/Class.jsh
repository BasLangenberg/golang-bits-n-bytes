public class TestClass {
    private String field1;

    public TestClass() {}

    public TestClass(String field1) {
        this.field1 = field1;
    }

    public String getField1() {
        return field1;
    }

    public void setField1(String field1) {
        this.field1 = field1;
    }

    @Override
    public boolean equals(Object o) {
        var tc = (TestClass) o;
        return tc.getField1().equals(this.getField1());
    }
}

var tc1 = new TestClass();
tc1.setField1("waarde1");
System.out.println(tc1.getField1());

var tc2 = new TestClass("waarde2");
System.out.println(tc2.getField1());

var tc3 = new TestClass("waarde1");
System.out.println("" + (tc1 == tc3) + " | " + tc1.equals(tc3))
