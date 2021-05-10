package top.azzgo.datastructure

import junit.framework.TestCase

object RationalTest : TestCase() {
    fun testPlus() {
        val a = Rational(4, 6)
        val b = Rational(1, 2)
        val sum = a.plus(b)
        assertEquals(7, sum.numerator)
        assertEquals(6, sum.denominator)
    }

    fun testMinus() {
        val a = Rational(5, 6)
        val b = Rational(14, 16)
        val result = a.minus(b)
        assertEquals(-1, result.numerator)
        assertEquals(24, result.denominator)
    }

    fun testTimes() {
        val a = Rational(10, 8)
        val b = Rational(14, 16)
        val result = a.times(b)
        assertEquals(35, result.numerator)
        assertEquals(32, result.denominator)
    }

    fun testDivides() {
        val a = Rational(10, 8)
        val b = Rational(14, 16)
        val result = a.divides(b)
        assertEquals(10, result.numerator)
        assertEquals(7, result.denominator)
    }

    fun testEquals() {
        val a = Rational(4, 8)
        val b = Rational(3, 6)
        assertTrue(a.equals(a))
        assertTrue(a.equals(b))
    }
}