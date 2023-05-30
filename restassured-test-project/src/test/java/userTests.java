import io.restassured.RestAssured;
import io.restassured.http.ContentType;
import io.restassured.response.Response;
import io.restassured.specification.RequestSpecification;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;
import static io.restassured.RestAssured.requestSpecification;
public class userTests {

    private static RequestSpecification request;
    public static Response response;
    public static String token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJzZkFhNlFwLWozSDhrSG5UbHE3cDZIdi1WYXdHVXBLZ25tTGhHUWZNaFlJIn0.eyJleHAiOjE2NzA1NTU3MTMsImlhdCI6MTY3MDU1NTQxMywianRpIjoiYjU5YmRmMmMtNGNlZS00YzkyLThiZTMtZmMwM2Y2ZTYxMzRkIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9tb25leS1ob3VzZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJmM2Q1ZWE4NS0xZTJjLTRkNGMtODllYi00YjQ1ZTUwNDA1ODYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ1c2Vycy1hcGkiLCJzZXNzaW9uX3N0YXRlIjoiOTYyMDg1NzAtODc5NC00MzhjLTgyYTctMzlkZjNkOGRjYjZlIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJkZWZhdWx0LXJvbGVzLW1vbmV5LWhvdXNlIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiI5NjIwODU3MC04Nzk0LTQzOGMtODJhNy0zOWRmM2Q4ZGNiNmUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6IlBydWViYURhbiBBdXRoIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidXNlcjEyM0BleGFtcGxlLmNvbSIsImdpdmVuX25hbWUiOiJQcnVlYmFEYW4iLCJmYW1pbHlfbmFtZSI6IkF1dGgiLCJlbWFpbCI6InVzZXIxMjNAZXhhbXBsZS5jb20ifQ.AyNiixY3Rd7Z-DCKjE3Sdnsf-LR04H9rqL57jWalgX2h5LfMfWud_YJigXGTrhJLbvEDADLrB0KEwYFQUlOPbwEUNMDzGiff2eWza27F5zl-UbeqZQfU1S0TmFc8uj0AMRgdoLSi2nX0txGQSH95aaxYYa19mKOMlbtBhdkyf6BG3lh6mH6vdkWrio1DLg8YYbTyByXdGYeE_C4COuC47Qdo1tcCWO76sq0gxv2xysrD81XHAPvUoG4tOu_Fe9Du5K4GKALxSXCOmzr-7DiZ3LlgrPU-DKgygS_eNKtAKynxuvdYYBpI4qlHGLQ46KWtVlly_7FB0OtlzKh1R3f65Q";


    public String oauth2Payload = "{\n" +
            " \"email\": \"user123@example.com\",\n" +
            " \"password\": \"user123\" \n}";
    public String getAccessToken(String oauth2Payload) {
        return given()
                .contentType(ContentType.JSON)
                .body(oauth2Payload)
                .post("/token")
                .then().extract().response()
                .jsonPath().getString("access_token");
    }

    public void userAdminConfigSetup() {
        requestSpecification = given().auth().oauth2(getAccessToken(oauth2Payload))
                .header("Accept", ContentType.JSON.getAcceptHeader())
                .contentType(ContentType.JSON);
    }

    @BeforeAll
    public static void beforeAll(){
        RestAssured.baseURI ="http://localhost:8081/api/v1/";
        request = RestAssured.given();
    }

    @Test
    @DisplayName("Get - AllUsers - OK")
    public void allUsersOK(){
        given()
                .auth().oauth2(token)
                .when()
                .get("/users/")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Get - AllUsers - OK")
    public void allUsers(){
        userAdminConfigSetup();
        response = given(requestSpecification).
                get("/users/").
                then().extract().response();

        //Assertions.assertEquals(201, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - OK")
    public void createUser(){
        final String requestBody = "{\n" +
                "    \"first_name\": \"User-Rest11\",\n" +
                "    \"lastName\": \"Assured1\",\n" +
                "    \"dni\": \"222222211\",\n" +
                "    \"email\": \"UserRe1tAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }


    @Test
    @DisplayName("Post - User - Name Blank - Err")
    public void createUserNameBlank(){
        final String requestBody = "{\n" +
                "    \"name\": \"\",\n" +
                "    \"lastName\": \"Assured\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"UserRestAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - Without Name - Err")
    public void createUserWithoutName(){
        final String requestBody = "{\n" +
                "    \"lastName\": \"Assured\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"UserRestAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - LastName Blank - Err")
    public void createUserLastNameBlank(){
        final String requestBody = "{\n" +
                "    \"name\": \"\",\n" +
                "    \"lastName\": \"\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"UserRestAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - Without LastName - Err")
    public void createUserWithOutLastName(){
        final String requestBody = "{\n" +
                "    \"name\": \"\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"UserRestAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - DNI Blank - Err")
    public void createUserDNIBlank(){
        final String requestBody = "{\n" +
                "    \"name\": \"Daniel\",\n" +
                "    \"lastName\": \"\",\n" +
                "    \"dni\": \"\",\n" +
                "    \"email\": \"UserRestAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - WithOut DNI - Err")
    public void createUserWithOutDNI(){
        final String requestBody = "{\n" +
                "    \"name\": \"Daniel\",\n" +
                "    \"lastName\": \"Romero\",\n" +
                "    \"email\": \"UserRestAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - Email Blank - Err")
    public void createUserEmailBlank(){
        final String requestBody = "{\n" +
                "    \"name\": \"Daniel\",\n" +
                "    \"lastName\": \"Romero\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - WithOut Email - Err")
    public void createUserWithOutEmail(){
        final String requestBody = "{\n" +
                "    \"name\": \"DAn\",\n" +
                "    \"lastName\": \"Rod\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - Telephone Blank - Err")
    public void createUserTelephoneBlank(){
        final String requestBody = "{\n" +
                "    \"name\": \"Daniel\",\n" +
                "    \"lastName\": \"Romero\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"dsr@gmail.com\",\n" +
                "    \"telephone\": \"\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Post - User - WithOut Telephone - Err")
    public void createUserWithOutTelephone(){
        final String requestBody = "{\n" +
                "    \"name\": \"Daniel\",\n" +
                "    \"lastName\": \"Romero\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"dsr@gmail.com\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Get - UserByID - OK")
    public void userByID(){

        request
                .auth().oauth2(token)
                .when()
                .get("/users/18c684da-6708-4516-b490-b5373d27f62c")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Get - UserByID - ID inexistente - Err")
    public void userByIDNonExist(){

        request
                .auth().oauth2(token)
                .when()
                .get("/users/1000000")
                .then()
                .statusCode(400)
                .log().all();
    }

    @Test
    @DisplayName("Get - UserByID - Wrong ID - Err")
    public void userByIDWrong(){

        request
                .auth().oauth2(token)
                .when()
                .get("/users/g")
                .then()
                .statusCode(400)
                .log().all();
    }

    @Test
    @DisplayName("Get - UserByID - Negative ID - Err")
    public void userByNegativeID(){

        request
                .auth().oauth2(token)
                .when()
                .get("/users/-8")
                .then()
                .statusCode(400)
                .log().all();
    }

    @Test
    @DisplayName("Patch - User - Err . Falta aut")
    public void userPatchOK(){
        final String requestBody = "{\n" +
                "    \"name\": \"Prueba\",\n" +
                "    \"lastName\": \"Patch\",\n" +
                "    \"dni\": \"2222\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .patch("/users/1")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    };

    @Test
    @DisplayName("Patch - User - Inexistente Id - Err")
    public void userPatchInexId(){
        final String requestBody = "{\n" +
                "    \"name\": \"Pruebsa\",\n" +
                "    \"lastName\": \"Patsch\",\n" +
                "    \"dni\": \"222245\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .patch("/users/845")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    };

    @Test
    @DisplayName("Patch - User - Rango equivocado Id - Err")
    public void userPatchIdRangeOut(){
        final String requestBody = "{\n" +
                "    \"name\": \"Pruebsa\",\n" +
                "    \"lastName\": \"Patsch\",\n" +
                "    \"dni\": \"222245\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .patch("/users/-5")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    };

    @Test
    @DisplayName("Patch - User - Id Invalido - Err")
    public void userPatchInvalidId(){
        final String requestBody = "{\n" +
                "    \"name\": \"Pruebsa\",\n" +
                "    \"lastName\": \"Patsch\",\n" +
                "    \"dni\": \"222245\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .patch("/users/g")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    };

    @Test
    @DisplayName("Patch - User - WithOut Id - Err")
    public void userPatchWithOutId(){
        final String requestBody = "{\n" +
                "    \"name\": \"Pruebsa\",\n" +
                "    \"lastName\": \"Patsch\",\n" +
                "    \"dni\": \"222245\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .patch("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(404, response.statusCode());
    };

    @Test
    @DisplayName("Delete - User - Ok")
    public void deleteUserByID() {
        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .when()
                .delete("/users/6")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Delete - User By ID inexistente- Err")
    public void deleteUserByIDInexistente() {
        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .when()
                .delete("/users/8450")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Delete - User By ID erroneo- Err")
    public void deleteUserByIDErr() {
        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .when()
                .delete("/users/h")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }

    @Test
    @DisplayName("Delete - User By ID fuera de rango- Err")
    public void deleteUserByIDRangeOut() {
        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .when()
                .delete("/users/-8")
                .then()
                .extract().response();

        Assertions.assertEquals(400, response.statusCode());
    }


    @Test
    @DisplayName("Get - AllUsers - Err - Unauthorized")
    public void allUsersNoAuth(){

        request
                .when()
                .get("/users/")
                .then()
                .statusCode(401)
                .log().all();
    }

    @Test
    @DisplayName("Post - User - Err - Unauthorized")
    public void createUserNoAuth(){
        final String requestBody = "{\n" +
                "    \"first_name\": \"User-Rest\",\n" +
                "    \"lastName\": \"Assured\",\n" +
                "    \"dni\": \"222222222\",\n" +
                "    \"email\": \"UserRestAssured@gmail.com\",\n" +
                "    \"telephone\": \"5493585478963\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/users/")
                .then()
                .extract().response();

        Assertions.assertEquals(401, response.statusCode());
    }

    @Test
    @DisplayName("Get - UserByID - Err - Unauthorized")
    public void userByIDNoAuth(){

        request
                .when()
                .get("/users/1")
                .then()
                .statusCode(401)
                .log().all();
    }

    @Test
    @DisplayName("Patch - User - Err - Unauthorized")
    public void userPatchNoAuth(){
        final String requestBody = "{\n" +
                "    \"name\": \"Prueba\",\n" +
                "    \"lastName\": \"Patch\",\n" +
                "    \"dni\": \"2222\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .patch("/users/1")
                .then()
                .extract().response();

        Assertions.assertEquals(401, response.statusCode());
    };

    @Test
    @DisplayName("Delete - User - Err - Unauthorized")
    public void deleteUserByIDNoAuth() {
        Response response = given()
                .header("Content-type", "application/json")
                .when()
                .delete("/users/1")
                .then()
                .extract().response();

        Assertions.assertEquals(401, response.statusCode());
    }
}
