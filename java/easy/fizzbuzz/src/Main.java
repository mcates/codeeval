/**
 * Created by Michael on 9/15/2015.
 */
import java.io.*;

public class Main {
    public static void main(String[] args) throws IOException {
        File file = new File(args[0]);
        BufferedReader buffer = new BufferedReader(new FileReader(file));
        String line;
        while ((line = buffer.readLine()) != null) {
            line = line.trim();
            String[] parts = line.split(" ");
            int fizz = Integer.parseInt(parts[0]);
            int buzz = Integer.parseInt(parts[1]);
            int length = Integer.parseInt(parts[2]);
            System.out.println(FizzBuzz(fizz, buzz, length));
        }
    }

    public static String FizzBuzz(int f, int b, int l) {
        String return_string = "";
        for (int i = 1; i <= l; i++) {
            if (i%f == 0 && i%b == 0) {
                return_string += "FB ";
            } else if (i%f == 0) {
                return_string += "F ";
            } else if (i%b == 0) {
                return_string += "B ";
            } else {
                return_string += i + " ";
            }
        }
        return return_string;
    }
}