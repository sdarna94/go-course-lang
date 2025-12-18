public class patten {
    
    public static void main(String[] args) {
        int n = 5;

        for (int i = 1; i <= n; i++) {
            int spaces = n - i;
            int stars  = 2 * i - 1;
            System.out.println(" ".repeat(spaces) + "*".repeat(stars));
        }
    }

    
}
