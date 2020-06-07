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

    public void printMe() {
        System.out.println(field1);
    }

    public static void printThis(TestClass tc) {
        System.out.println(tc.getField1());
    }
}

var tc1 = new TestClass("waarde1");
TestClass.printThis(tc1);

var tc2 = new TestClass("waarde2");
tc2.printMe()
