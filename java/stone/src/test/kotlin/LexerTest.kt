import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.runBlocking
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals


internal class LexerTest {
    @Test
    fun givenACodeStringItWillParseAsTokenList() = runBlocking {
        val InputCode = """
           while i < 10 {
                sum = sum + i
                i=i+1
            }
            sum
        """
        val l = Lexer(InputCode.reader())


        val tokenList = flow<String> {
            var t: Token
            while ((l.read().also { t = it }) != Token.EOF) {
                emit(t.getText())
            }
        }.toList()

        assertEquals(
            tokenList,
            listOf(
                "\\n",
                "while",
                "i",
                "<",
                "10",
                "{",
                "\\n",
                "sum",
                "=",
                "sum",
                "+",
                "i",
                "\\n",
                "i",
                "=",
                "i",
                "+",
                "1",
                "\\n",
                "}",
                "\\n",
                "sum",
                "\\n",
                "\\n"
            )
        )
    }

}