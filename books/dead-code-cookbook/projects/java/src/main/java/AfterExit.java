public class AfterExit {
    public static void main(String[] args) {
        try {
            System.exit(0);
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }
}