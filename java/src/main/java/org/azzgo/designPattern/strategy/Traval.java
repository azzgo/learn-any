package org.azzgo.designPattern.strategy;


public  class Traval {
    private String target;

    Traval(String targetName) {
        this.target = targetName;
    }

    public void goToAirport(Vehicle vehicle) {
        System.out.println("=====================");
        System.out.println("ready to Airport.");
        vehicle.transition();
        System.out.println("We arrived the Airport.");
    }

    public void flyViaAirPlane() {
        System.out.println("=====================");
        System.out.println("ready to " + target + ".");
        Vehicle airplane = new AirPlane();
        airplane.transition();
        System.out.println("We has already arrived！the " + target + " !");
    }


    void lookAndPlaying() {
        System.out.println("=====================");
        System.out.println("We are traveling on " + target + ", and take photos.");
    }

    void goToHotel(Vehicle vehicle) {
        System.out.println("=====================");
        System.out.println("ready to hotel");
        vehicle.transition();
        System.out.println("We on hotel now!");
    }
}
