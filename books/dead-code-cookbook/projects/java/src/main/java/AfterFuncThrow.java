public class AfterFuncThrow {
    public static void main(String[] args) {
        try {
            f();
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }

    private static void f() {
        throw new RuntimeException();
    }
}