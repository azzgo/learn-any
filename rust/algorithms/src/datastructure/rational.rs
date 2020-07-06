struct Rational {
    pub numerator: i32,
    pub denominator: i32,
}

impl Rational {
    pub fn plus(&self, b: Rational) -> Rational {
        let new_numerator = self.numerator * b.denominator + self.denominator * b.numerator;
        let new_denominator = self.denominator * b.denominator;
        return Rational {
            numerator: new_numerator,
            denominator: new_denominator,
        };
    }

    pub fn minus(&self, b: Rational) -> Rational {
        let new_numerator = self.numerator * b.denominator - self.denominator * b.numerator;
        let new_denominator = self.denominator * b.denominator;
        return Rational {
            numerator: new_numerator,
            denominator: new_denominator,
        };
    }

    pub fn times(&self, b: Rational) -> Rational {
        let new_numerator = self.numerator * b.numerator;
        let new_denominator = self.denominator * b.denominator;
        return Rational {
            numerator: new_numerator,
            denominator: new_denominator,
        };
    }


    pub fn divides(&self, b: Rational) -> Rational {
        let new_numerator = self.numerator * b.denominator;
        let new_denominator = self.denominator * b.numerator;
        return Rational {
            numerator: new_numerator,
            denominator: new_denominator,
        };
    }
}

impl Rational {
    pub fn new(numerator: i32, denominator: i32) -> Rational {
        let cd = gcd(numerator.abs(), denominator.abs());

        if cd != 1 {
            return Rational {
                numerator: numerator / cd,
                denominator: denominator / cd,
            };
        } else {
            return Rational {
                numerator: numerator,
                denominator: denominator,
            };
        }
    }
}

fn gcd(a: i32, b: i32) -> i32 {
    let mut devider = b;
    let mut num = a;
    let mut remainder = num % devider;

    while remainder != 0 {
        num = devider;
        devider = remainder;
        remainder = num % devider;
    }

    return devider;
}

#[cfg(test)]
mod tests {
    use super::Rational;

    #[test]
    fn it_should_be_1_2() {
        let rational1 = Rational::new(3, 6);

        assert_eq!(rational1.numerator, 1);
    }

    #[test]
    fn it_plus_should_be_5_6() {
        let rational1 = Rational::new(1, 2);
        let rational2 = Rational::new(2, 6);

        let sum = rational1.plus(rational2);

        assert_eq!(sum.numerator, 5);
        assert_eq!(sum.denominator, 6);
    }

    #[test]
    fn it_minus_should_be_1_6() {
        let rational1 = Rational::new(1, 2);
        let rational2 = Rational::new(2, 6);

        let result = rational1.minus(rational2);

        assert_eq!(result.numerator, 1);
        assert_eq!(result.denominator, 6);
    }

    #[test]
    fn it_times_should_be_1_6() {
        let rational1 = Rational::new(1, 2);
        let rational2 = Rational::new(2, 6);

        let result = rational1.times(rational2);

        assert_eq!(result.numerator, 1);
        assert_eq!(result.denominator, 6);
    }

    #[test]
    fn it_divides_should_be_3_2() {
        let rational1 = Rational::new(1, 2);
        let rational2 = Rational::new(2, 6);

        let result = rational1.divides(rational2);

        assert_eq!(result.numerator, 3);
        assert_eq!(result.denominator, 2);
    }
}
