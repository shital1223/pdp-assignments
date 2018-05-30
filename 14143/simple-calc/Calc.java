import java.util.*;
import java.util.Scanner;
public class Calc {
    public static void main(String[] args)
    {
        System.out.println("Enter first and second number:");
        Scanner inp= new Scanner(System.in);
        long num1,num2;
        num1 = inp.nextLong();
        num2 = inp.nextLong();
        long ans;
        System.out.println("Enter your selection: 1 for Addition, 2 for substraction,3 for Multiplication,4 for division and 5 for mod:");
        int choose;
        choose = inp.nextInt();
        switch (choose){
        case 1:
            System.out.println(add( num1,num2));
            break;
        case 2:
            System.out.println(sub( num1,num2));
            break;      
        case 3:
            System.out.println(mult( num1,num2));
            break;
        case 4:
            System.out.println(div( num1,num2));
            break;
        case 5:
            System.out.println(mod( num1,num2));
            break;
            default:
                System.out.println("Illigal Operation");

        }

    }
    public static long add(long x, long y)
    {
        long result = x + y;
        return result;
    }
    public static long sub(long x, long y)
    {
        long result = x-y;
        return result;
    }
    public static long mult(long x, long y)
    {
        long result = x*y;
        return result;
    }
    public static long div(long x, long y)
    {
        long result = x/y;
        return result;
    }
    public static long mod(long x, long y)
    {
        long result = x%y;
        return result;
    }

}
