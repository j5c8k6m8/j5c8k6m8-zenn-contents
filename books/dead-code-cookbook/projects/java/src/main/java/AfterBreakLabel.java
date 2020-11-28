public class AfterBreakLabel {
    public static void main(String[] args) {
        l:{
            while(true) {
                while(true) {
                    break l;
                }
                System.out.println("Am I dead?");
            }
        }
    }
}