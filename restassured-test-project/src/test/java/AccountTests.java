import io.restassured.RestAssured;
import io.restassured.response.Response;
import io.restassured.specification.RequestSpecification;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;

public class AccountTests {
    private static RequestSpecification request;
    public static String token="eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJzZkFhNlFwLWozSDhrSG5UbHE3cDZIdi1WYXdHVXBLZ25tTGhHUWZNaFlJIn0.eyJleHAiOjE2NzA1NTcwODgsImlhdCI6MTY3MDU1Njc4OCwianRpIjoiMDBiNjZjZmUtODdlMS00M2IyLTg1ZTUtNjQ1ZjNhYmNiOGEyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9tb25leS1ob3VzZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJmM2Q1ZWE4NS0xZTJjLTRkNGMtODllYi00YjQ1ZTUwNDA1ODYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ1c2Vycy1hcGkiLCJzZXNzaW9uX3N0YXRlIjoiY2M2YTBlMGYtMjZmYi00Y2U3LTlmMWEtODVlMWUyYzA5MDM1IiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJkZWZhdWx0LXJvbGVzLW1vbmV5LWhvdXNlIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiJjYzZhMGUwZi0yNmZiLTRjZTctOWYxYS04NWUxZTJjMDkwMzUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6IlBydWViYURhbiBBdXRoIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidXNlcjEyM0BleGFtcGxlLmNvbSIsImdpdmVuX25hbWUiOiJQcnVlYmFEYW4iLCJmYW1pbHlfbmFtZSI6IkF1dGgiLCJlbWFpbCI6InVzZXIxMjNAZXhhbXBsZS5jb20ifQ.r9ULR4NzeHfMGvB8O9f5rNKpwlrPAV_p3JJ8FWC6WyMKplHWgnoRT5PH4iz_8P52AjhcjPqiru4Qzua0TswvMc-jhFWgA62DBjmqAttlJ02duEYC_A0-VdF-qmkVAmglVT23GC1psbu2SzFxY1-qRygODJbjHIeepFshbl26-yDiD3RhL7hU9wCD5LW0syqPeCR9REm0-kcpl2UzKeCP0QHidI3eKOVaqgb-ONSYNmeuKplydwPrrt8B-Yt6G8437Bdxz5smd4Q2Kyyn7RrMg8PntOWw1rX5vgmM73UBdqrW7KZcU0iDmhH98E-SDKqtq0zSh0O5PsZLLwjdO-r5qg";

    @BeforeAll
    public static void beforeAll(){
        RestAssured.baseURI ="http://localhost:8083/api/v1";
        request = given();
    }

    @Test
    @DisplayName("Get - AllAccounts - OK")
    public void allAccounts(){

        request
                .auth().oauth2(token)
                .when()
                .get("/accounts/")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Post - Account - OK")
    public void createAccount(){
        final String requestBody = "{\n" +
                "    \"amount\": 2349.00,\n" +
                "    \"user_id\": 1 \n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/accounts/")
                .then()
                .extract().response();

        Assertions.assertEquals(201, response.statusCode());
    }

    @Test
    @DisplayName("Get - AccountByID - OK")
    public void acountByID(){

        request
                .auth().oauth2(token)
                .when()
                .get("/accounts/1")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Get - AllAccounts - Err - Unauthorized")
    public void allAccountsNoAuth(){

        request
                .when()
                .get("/accounts/")
                .then()
                .statusCode(404)
                .log().all();
    }

    @Test
    @DisplayName("Post - Account - Err - Unauthorized")
    public void createAccountNoAuth(){
        final String requestBody = "{\n" +
                "    \"amount\": 2349.00,\n" +
                "    \"user_id\": 19 \n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/accounts/")
                .then()
                .extract().response();

        Assertions.assertEquals(401, response.statusCode());
    }

    @Test
    @DisplayName("Get - AccountByID - Err - Unauthorized")
    public void acountByIDNoAuth(){

        request
                .when()
                .get("/accounts/1")
                .then()
                .statusCode(403)
                .log().all();
    }
}
