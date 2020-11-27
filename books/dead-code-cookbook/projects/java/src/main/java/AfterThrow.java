public class AfterThrow {
    public static void main(String[] args) {
        try {
            throw new RuntimeException();
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }
}