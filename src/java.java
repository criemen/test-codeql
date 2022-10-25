import java.security.SecureRandom;

public class java {
	public static void main(String[] args) {
		SecureRandom random = new SecureRandom();
		random.setSeed(0L); // java/predictable-seed
		random.nextBytes(new byte[32]);
	}
}
