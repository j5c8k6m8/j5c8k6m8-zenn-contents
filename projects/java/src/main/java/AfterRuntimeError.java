public class AfterRuntimeError {
    public static void main(String[] args) {
        try {
            Integer a = null, b = null;
            int x = a + b;
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }
}