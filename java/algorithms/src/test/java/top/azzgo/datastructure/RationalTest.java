package top.azzgo.datastructure;

import junit.framework.TestCase;

public class RationalTest extends TestCase {
    public static void testPlus() {
        Rational a = new Rational(4, 6);
        Rational b = new Rational(1, 2);

        Rational sum = a.plus(b);


        assertEquals(7, sum.getNumerator());
        assertEquals(6, sum.getDenominator());
    }

    public static void testMinus() {
        Rational a = new Rational(5, 6);
        Rational b = new Rational(14, 16);

        Rational result = a.minus(b);

        assertEquals(-1, result.getNumerator());
        assertEquals(24, result.getDenominator());
    }

    public static void testTimes() {
        Rational a = new Rational(10, 8);
        Rational b = new Rational(14, 16);

        Rational result = a.times(b);

        assertEquals(35, result.getNumerator());
        assertEquals(32, result.getDenominator());
    }

    public static void testDivides() {

        Rational a = new Rational(10, 8);
        Rational b = new Rational(14, 16);

        Rational result = a.divides(b);

        assertEquals(10, result.getNumerator());
        assertEquals(7, result.getDenominator());
    }

    public static void testEquals() {
        Rational a = new Rational(4, 8);
        Rational b = new Rational(3, 6);

        assertTrue(a.equals(a));
        assertTrue(a.equals(b));
    }
}
