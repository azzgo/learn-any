package org.azzgo.designPattern.state;

import org.azzgo.designPattern.state.states.GreenSignal;
import org.junit.Test;

public class StateTest {
    @Test
    public void shouldWork() {
        TrafficLight jorgeAvenueTrafficLight = new JorgeAvenueTrafficLight();

        jorgeAvenueTrafficLight.setSignal(new GreenSignal());

        jorgeAvenueTrafficLight.nextSignal();

        System.out.println("=========");

        jorgeAvenueTrafficLight.nextSignal();

        System.out.println("=========");

        jorgeAvenueTrafficLight.nextSignal();
    }
}
