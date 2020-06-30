package top.azzgo.datastructure;

public class Rational {
    private final int numerator;
    private final int denominator;

    private static int gcd(int a, int b) {
        int divider = b;
        int num = a;

        int remainer = num % divider;

        while (remainer != 0) {
            num = divider;
            divider = remainer;
            remainer = num % divider;
        }


        return divider;
    }

    Rational(int numerator, int denominator) {

        int cd = gcd(Math.abs(numerator), Math.abs(denominator));

        if (cd != 1) {
            this.numerator = numerator / cd;
            this.denominator = denominator / cd;
        } else {
            this.numerator = numerator;
            this.denominator = denominator;
        }

    }


   public Rational plus(Rational b) {
        int newNumerator = this.numerator * b.denominator + this.denominator * b.numerator;
        int newDemoninator = this.denominator * b.denominator;

        return new Rational(newNumerator, newDemoninator);
    }


    public Rational minus(Rational b) {
        int newNumerator = this.numerator * b.denominator - this.denominator * b.numerator;
        int newDemoninator = this.denominator * b.denominator;

        return new Rational(newNumerator, newDemoninator);
    }


    public Rational times(Rational b) {

        int newNumerator = this.numerator * b.numerator;
        int newDemoninator = this.denominator * b.denominator;

        return new Rational(newNumerator, newDemoninator);
    }

    public Rational divides(Rational b) {

        int newNumerator = this.numerator * b.denominator;
        int newDemoninator = this.denominator * b.numerator;

        return new Rational(newNumerator, newDemoninator);
    }


    public int getNumerator() {
        return numerator;
    }

    public int getDenominator() {
        return denominator;
    }

    @Override
    public String toString() {
        return this.numerator + "/" + this.denominator;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }

        if (obj.getClass() != Rational.class) {
            return false;
        }
        Rational b = (Rational) obj;

        return this.numerator == b.numerator && this.denominator == b.denominator;
    }
}
