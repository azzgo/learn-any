package org.azzgo.designPattern.strategy;

import org.junit.Assert;
import org.junit.Test;

public class StrategyTest {

    @Test
    public void shouldWork() {
        Traval travalToJapan = new Traval("🇯🇵");

        travalToJapan.goToAirport(new Bus());

        travalToJapan.flyViaAirPlane();

        travalToJapan.lookAndPlaying();

        travalToJapan.goToHotel(new Walk());

    }
}
