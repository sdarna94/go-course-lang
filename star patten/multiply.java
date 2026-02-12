// User function Template for Java
import java.util.*; 

class multiply {
    
    public ArrayList<Integer> getTable(int n) {
        ArrayList<Integer> res = new ArrayList<>(10);
        for (int i = 1; i <= 10; i++) {
            res.add(n * i);    // n*1 ... n*10 
        }
        return res;
    }
            
         public static void main(String[] args) {
             
       Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        sc.close();
        
        multiply ob = new multiply();
        ArrayList<Integer> ans = ob.getTable(n);   // ArrayList version

        for (int x : ans) {
            System.out.print(x + " ");
        }
    }
}
