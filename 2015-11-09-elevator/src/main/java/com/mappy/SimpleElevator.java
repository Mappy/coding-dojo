package com.mappy;

import java.util.List;

import static com.google.common.collect.Lists.*;

public class SimpleElevator implements Elevator {
    private final List<String> commands = newArrayList(
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "DOWN",
            "DOWN",
            "DOWN",
            "DOWN",
            "DOWN");
    private int i = -1;
    private String required;

    @Override
    public void reset() {}

    @Override
    public void userHasEntered() {}

    @Override
    public void userHasExited() {}

    @Override
    public void call(String atFloor, String to) {
        required = atFloor;
    }

    @Override
    public void go(String floorToGo) {}

    @Override
    public String nextCommand() {
        i = (i + 1) % commands.size();
        return commands.get(i);
    }
}
