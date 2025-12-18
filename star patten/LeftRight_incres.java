class LeftRight_incres {

    public static void main(String[] args) { // TODO Auto-generated method stub
       int n =5; //number of rows 
       for(int i=1;i<=n;i++){    //outer loop for number of rows
        for (int j=i;j<=n;j++){  //inner loop for spaces
            System.out.print(" "); //print space
        }
        for(int j=1;j<=i;j++){  //inner loop for stars
                System.out.print("*"); //print star

            }

        System.out.println(""); //move to next line after each row

       }

 }
}