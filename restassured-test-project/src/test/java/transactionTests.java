import io.restassured.RestAssured;
import io.restassured.response.Response;
import io.restassured.specification.RequestSpecification;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;

public class transactionTests {
    private static RequestSpecification request;
    public static String token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJzZkFhNlFwLWozSDhrSG5UbHE3cDZIdi1WYXdHVXBLZ25tTGhHUWZNaFlJIn0.eyJleHAiOjE2NjkyNjY1MTksImlhdCI6MTY2OTI2NjIxOSwianRpIjoiODBkNGNhNjctOGFhZi00ZDMzLWEyZDQtYjU1ODlhMmMyMGVjIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9tb25leS1ob3VzZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyNmQzZTE2OC05ZTdmLTQ5M2UtOTlkNy1iOGI0N2E0M2VlOWQiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ1c2Vycy1hcGkiLCJzZXNzaW9uX3N0YXRlIjoiYWEzMDA4ZjUtNDkwZS00YjZkLWFkMjEtZDIxZDVkMWJjOWNlIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJkZWZhdWx0LXJvbGVzLW1vbmV5LWhvdXNlIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiJhYTMwMDhmNS00OTBlLTRiNmQtYWQyMS1kMjFkNWQxYmM5Y2UiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6IkRhbiBSb20iLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJ1c2VyQGV4YW1wbGUuY29tIiwiZ2l2ZW5fbmFtZSI6IkRhbiIsImZhbWlseV9uYW1lIjoiUm9tIiwiZW1haWwiOiJ1c2VyQGV4YW1wbGUuY29tIn0.tDHO6iY7OoWrawDGIQp7Uod13BcYG7QFu-HXy_dWj12bq-q83vb-9Nmm3_eKg05ojBRIoK0_Gm70xzprUXhwU5OguVDDhc7mbaVN6zkhsEVqAhHsZ8VMw6oeu1uQANmf6J03uVVbW0vs8jj_exgYyeHXHUzbi2QOhYmWKK_PKUaNrok73huY1QqkjiaZY59ZlSK47oCZXmDCQi-k9mNx0ZBryRrU58KihoL_6fqBLIi1o3FwIL9Lf40dD_yk9iNZFASKbVMEdJ7iMsg5ZKQ86EqvevS95twv6QtEXLP0ZNKrtXfC4RoP8FAoAiSsyh1NTBjykJK_44toSzyzVsLuKg";


    @BeforeAll
    public static void beforeAll(){
        RestAssured.baseURI ="http://localhost:8083/api/v1";
        request = given();
    }

    @Test
    @DisplayName("Get - AllTransactions - OK")
    public void allCards(){

        request
                .auth().oauth2(token)
                .when()
                .get("/transactions/")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Post - Transaction - OK")
    public void createTransaction(){
        final String requestBody = "{\n" +
                "    \"amount\": 10,\n" +
                "    \"transaction_date\": \"2022-11-11 13:23:44\",\n" +
                "    \"description\": \"Varios\",\n" +
                "    \"origin_cvu\": \"4389564134587078096588\",\n" +
                "    \"destination_cvu\": \"1316147578510646806002\",\n" +
                "    \"account_id\": 1,\n" +
                "    \"transaction_type\": \"egreso\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/transactions")
                .then()
                .extract().response();

        Assertions.assertEquals(201, response.statusCode());
    }

    @Test
    @DisplayName("Get - AllTransactions - Err - Unauthorized")
    public void allCardsNoAuth(){

        request
                .when()
                .get("/transactions/")
                .then()
                .statusCode(404)
                .log().all();
    }

    @Test
    @DisplayName("Post - Transaction - OK")
    public void createTransactionNoAuth(){
        final String requestBody = "{\n" +
                "    \"amount\": 10,\n" +
                "    \"transaction_date\": \"2022-11-11 13:23:44\",\n" +
                "    \"description\": \"Varios\",\n" +
                "    \"origin_cvu\": \"4389564134587078096588\",\n" +
                "    \"destination_cvu\": \"1316147578510646806002\",\n" +
                "    \"account_id\": 1,\n" +
                "    \"transaction_type\": \"egreso\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/transactions")
                .then()
                .extract().response();

        Assertions.assertEquals(307, response.statusCode());
    }
}
