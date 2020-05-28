public interface TestInterface {

    static TestInterface getInstance() { return new TestClassOne(); }
}

public class TestClassOne implements TestInterface {

}

public class TestClassTwo {

}

var ti1 = TestInterface.getInstance();
System.out.println(ti1 instanceof TestClassOne);
System.out.println(ti1 instanceof TestClassTwo);
