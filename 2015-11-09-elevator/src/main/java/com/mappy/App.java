package com.mappy;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import static spark.Spark.*;

public class App {
    private static final Logger LOGGER = LoggerFactory.getLogger(App.class);

    public static void main(String[] args) {
        Elevator elevator = new SimpleElevator();
        get("/reset", (req, res) -> {
            LOGGER.info("reset(cause: {})", req.queryParams("cause"));
            elevator.reset();
            return res;
        });
        get("/call", (req, res) -> {
            String atFloor = req.queryParams("atFloor");
            String to = req.queryParams("to");
            LOGGER.info("call(atFloor: {}, to: {})", atFloor, to);
            elevator.call(atFloor, to);
            return res;
        });
        get("/userHasEntered", (req, res) -> {
            LOGGER.info("userHasEntered");
            elevator.userHasEntered();
            return res;
        });
        get("/userHasExited", (req, res) -> {
            LOGGER.info("userHasExited");
            elevator.userHasExited();
            return res;
        });
        get("/go", (req, res) -> {
            String floorToGo = req.queryParams("floorToGo");
            LOGGER.info("go(floorToGo: {})", floorToGo);
            elevator.go(floorToGo);
            return res;
        });
        get("/nextCommand", (req, res) -> {
            String nextCommand = elevator.nextCommand();
            LOGGER.info("nextCommand = {}", nextCommand);
            return nextCommand;
        });
    }
}