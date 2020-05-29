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

    public void giveMeException() throws Exception {
        throw new Exception("here u go");
    }

    public void giveMeRuntimeException() {
        throw new RuntimeException("whoops");
    }
}

var tc1 = new TestClass("waarde1");
try {
    tc1.giveMeException();
} catch (Exception ex) {
    ex.printStackTrace();
}

var tc2 = new TestClass("waarde2");
try {
    tc2.giveMeRuntimeException();
} catch (Exception ex) {
    ex.printStackTrace();
}
