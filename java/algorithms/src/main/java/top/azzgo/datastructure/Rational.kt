package top.azzgo.datastructure

import top.azzgo.datastructure.Rational

class Rational internal constructor(numerator: Int, denominator: Int) {
    var numerator = 0
    var denominator = 0
    operator fun plus(b: Rational): Rational {
        val newNumerator = numerator * b.denominator + denominator * b.numerator
        val newDemoninator = denominator * b.denominator
        return Rational(newNumerator, newDemoninator)
    }

    operator fun minus(b: Rational): Rational {
        val newNumerator = numerator * b.denominator - denominator * b.numerator
        val newDemoninator = denominator * b.denominator
        return Rational(newNumerator, newDemoninator)
    }

    operator fun times(b: Rational): Rational {
        val newNumerator = numerator * b.numerator
        val newDemoninator = denominator * b.denominator
        return Rational(newNumerator, newDemoninator)
    }

    fun divides(b: Rational): Rational {
        val newNumerator = numerator * b.denominator
        val newDemoninator = denominator * b.numerator
        return Rational(newNumerator, newDemoninator)
    }

    override fun toString(): String {
        return numerator.toString() + "/" + denominator
    }

    override fun equals(obj: Any?): Boolean {
        if (this === obj) {
            return true
        }
        if (obj!!.javaClass != Rational::class.java) {
            return false
        }
        val b = obj as Rational?
        return numerator == b!!.numerator && denominator == b!!.denominator
    }

    companion object {
        private fun gcd(a: Int, b: Int): Int {
            var divider = b
            var num = a
            var remainer = num % divider
            while (remainer != 0) {
                num = divider
                divider = remainer
                remainer = num % divider
            }
            return divider
        }
    }

    init {
        assert(numerator < Int.MAX_VALUE)
        assert(numerator > Int.MIN_VALUE)
        assert(denominator < Int.MAX_VALUE)
        assert(denominator > Int.MIN_VALUE)
        val cd = gcd(Math.abs(numerator), Math.abs(denominator))
        if (cd != 1) {
            this.numerator = numerator / cd
            this.denominator = denominator / cd
        } else {
            this.numerator = numerator
            this.denominator = denominator
        }
    }
}