

import io.restassured.RestAssured;
import io.restassured.response.Response;
import io.restassured.specification.RequestSpecification;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.Random;

import static io.restassured.RestAssured.given;

public class authTest {

    public static int generateRandomIntIntRange(int min, int max) {
        Random r = new Random();
        return r.nextInt((max - min) + 1) + min;
    }
    private static RequestSpecification request;
    private String userNumber = String.valueOf(generateRandomIntIntRange(1,1000));


    @BeforeAll
    public static void beforeAll(){
        RestAssured.baseURI ="http://localhost:8081/api/v1/auth";
        request = given();
    }

    @Test
    @DisplayName("Post - Register - OK")
    public void Register(){
        final String requestBody = "{\n" +
                "    \"email\": \"user" + userNumber + "@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"PruebaRest" + userNumber + "\",\n" +
                "    \"lastName\": \"Assured" + userNumber + "\",\n" +
                "    \"dni\": \"" + userNumber + userNumber + "\",\n" +
                "    \"telephone\": \"" + userNumber + userNumber + "3\"\n" +
                "}";


        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(201, response.statusCode());
    }

    @Test
    @DisplayName("Post - Login - OK")
    public void Login(){
        final String requestBody = "{\n" +
                "    \"email\": \"user123@example.com\",\n" +
                "    \"password\": \"user123\"\n"+
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/login")
                .then()
                .extract().response();

        Assertions.assertEquals(200, response.statusCode());
    }

/*
    @Test
    @DisplayName("Post - Register - Name Blank - Err")
    public void RegisterNameBlank(){
        final String requestBody = "{\n" +
                "    \"email\": \"pruebath@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"\",\n" +
                "    \"lastName\": \"Autoh\",\n" +
                "    \"dni\": \"00.000.000\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Without Name - Err")
    public void RegisterWithoutName(){
        final String requestBody = "{\n" +
                "    \"email\": \"pruebath@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"lastName\": \"Autoh\",\n" +
                "    \"dni\": \"00.000.000\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Last Name Blank - Err")
    public void RegisterLastNameBlank(){
        final String requestBody = "{\n" +
                "    \"email\": \"pruebath@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"lastName\": \"\",\n" +
                "    \"dni\": \"00.000.000\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Without Last Name - Err")
    public void RegisterWithoutLastName(){
        final String requestBody = "{\n" +
                "    \"email\": \"pruebath@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"dni\": \"00.000.000\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - DNI Blank - Err")
    public void RegisterDNIBlank(){
        final String requestBody = "{\n" +
                "    \"email\": \"pruebath@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"lastName\": \"dfsewf\",\n" +
                "    \"dni\": \"\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Without DNI - Err")
    public void RegisterWithoutDNI(){
        final String requestBody = "{\n" +
                "    \"email\": \"pruebath@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"lastName\": \"dfsewf\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Email Blank - Err")
    public void RegisterEmailBlank(){
        final String requestBody = "{\n" +
                "    \"email\": \"\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"lastName\": \"dfsewf\",\n" +
                "    \"dni\": \"123456\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Without Email - Err")
    public void RegisterWithoutEmail(){
        final String requestBody = "{\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"lastName\": \"dfsewf\",\n" +
                "    \"dni\": \"123456\",\n" +
                "    \"telephone\": \"5485478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Telephone Blank - Err")
    public void RegisterTelephoneBlank(){
        final String requestBody = "{\n" +
                "    \"email\": \"prueba@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"lastName\": \"dfsewf\",\n" +
                "    \"dni\": \"123456\",\n" +
                "    \"telephone\": \"\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - Register - Without Telephone - Err")
    public void RegisterWithoutTelephone(){
        final String requestBody = "{\n" +
                "    \"email\": \"prueba@example.com\",\n" +
                "    \"password\": \"user123\",\n" +
                "    \"firstName\": \"DAaaa\",\n" +
                "    \"lastName\": \"dfsewf\",\n" +
                "    \"dni\": \"123456\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/register")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }



    @Test
    @DisplayName("Post - Login - Email erroneo - Err")
    public void LoginFail(){
        final String requestBody = "{\n" +
                "    \"email\": \"usera0a73695-abef4c5b3664f@example.com\",\n" +
                "    \"password\": \"user123\"\n"+
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/login")
                .then()
                .extract().response();

        Assertions.assertEquals(404, response.statusCode());
    }*/

}
