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

    public static void changeValueInt(int i) {
        i = 43;
    }

    public static void changeValueTestClass(TestClass tc1) {
        tc1.setField1("changed");
    }

    public static void changeReference(TestClass tc1) {
        tc1 = new TestClass("reference");
        tc1.setField1("changed reference");
    }
}

var i1 = 42;
TestClass.changeValueInt(i1);
System.out.println(42);

var tc1 = new TestClass("original");
TestClass.changeValueTestClass(tc1);
System.out.println(tc1.getField1());

var tc2 = new TestClass("copy");
TestClass.changeReference(tc2);
System.out.println(tc2.getField1());
