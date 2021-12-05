package top.azzgo.datastructure

import junit.framework.Assert
import junit.framework.TestCase
import kotlin.math.sqrt

class CollectionTest : TestCase() {

    fun testBullsStatsCalculator() {
        val numbers: Bag<Double> = Bag()

        // prepare data
        numbers.add(100.00)
        numbers.add(99.00)
        numbers.add(101.00)
        numbers.add(120.00)
        numbers.add(98.00)
        numbers.add(107.00)
        numbers.add(109.00)
        numbers.add(81.00)
        numbers.add(101.00)
        numbers.add(90.00)

        val N = numbers.size()

        // 和
        var sum = 0.0
        for (x in numbers) {
            sum += x
        }

        // 平均数
        val mean = sum / N

        sum = 0.0
        for (x in numbers) {
            sum += (x - mean) * (x - mean)
        }

        val std = sqrt(sum / N - 1)

        Assert.assertEquals(std, 10.51)
        Assert.assertEquals(mean, 100.60)
    }

    fun testQueue() {
        val testData = listOf("to", "be", "or", "not", "to", "-", "be", "-", "-", "that", "-", "-", "-", "is")
        val quene = Quene<String>()
        for (data in testData) {
            if (data != "-") {
                quene.enqueue(data)
            }
            else if (!quene.isEmpty()) {
                print(quene.dequeue() + " ")
            }
        }

        Assert.assertEquals(quene.size(), 2);
    }

    fun testStack() {
//       val stack: Stack<String> = ResizingStack<String>();
        val stack: Stack<String> = LinkStack<String>()

        stack.push("1")
        stack.push("2")
        stack.push("3")

        Assert.assertEquals(stack.size(), 3)
        Assert.assertEquals(stack.pop(), "3")
        Assert.assertEquals(stack.pop(), "2")
        Assert.assertEquals(stack.pop(), "1")
    }

}