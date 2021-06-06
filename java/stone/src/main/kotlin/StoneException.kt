import java.lang.RuntimeException

class StoneException(m: String, t: ASTree?): RuntimeException("$m ${t?.location}") {
    constructor(m: String) : this(m, null)
}