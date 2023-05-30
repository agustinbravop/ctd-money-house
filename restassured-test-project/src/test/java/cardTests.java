import io.restassured.RestAssured;
import io.restassured.response.Response;
import io.restassured.specification.RequestSpecification;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;

public class cardTests {
    private static RequestSpecification request;
    public static String token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJzZkFhNlFwLWozSDhrSG5UbHE3cDZIdi1WYXdHVXBLZ25tTGhHUWZNaFlJIn0.eyJleHAiOjE2NjkyNjY4MTIsImlhdCI6MTY2OTI2NjUxMiwianRpIjoiZmRkNzhlMTEtZDA1Mi00MGY4LTgwOWQtODMyZDE4ZDFjMDc2IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9tb25leS1ob3VzZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyNmQzZTE2OC05ZTdmLTQ5M2UtOTlkNy1iOGI0N2E0M2VlOWQiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ1c2Vycy1hcGkiLCJzZXNzaW9uX3N0YXRlIjoiZTNjMDM2NGMtNzliMi00ZTU2LWFmOWMtMzQ1MjE3OGNiOWYwIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJkZWZhdWx0LXJvbGVzLW1vbmV5LWhvdXNlIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiJlM2MwMzY0Yy03OWIyLTRlNTYtYWY5Yy0zNDUyMTc4Y2I5ZjAiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6IkRhbiBSb20iLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJ1c2VyQGV4YW1wbGUuY29tIiwiZ2l2ZW5fbmFtZSI6IkRhbiIsImZhbWlseV9uYW1lIjoiUm9tIiwiZW1haWwiOiJ1c2VyQGV4YW1wbGUuY29tIn0.dcqmlYK_vbS-UQx6nZvNfROVURTKpy-G5K8BtAXOdK4oUHtUyTDbSPeCzJOTawhXhgSFROKLvyvdQpJClCqFQ7InEyLRY-qWqwO5hnRAbFfjaX9KptmHJ7X7jYDazxwGHgEcn48c2IMb7jr_LSakof0JBrVlWumIXMHJGYebCguN47C-TXKmEPpCK46L9w_RBSt2l6jnDivoQxtZMoiJ6Q7xg1QYS8yizqYvSq-M7PbYh_3myFqizysArocgijHcbV32o5pWGlzVhpJbS1yIt-HurPeJTl0yKMkRVCUobF2N5CgCWbglTWthueTT5TqQxEfI7WgJOvGRp8Weu6VFpQ";


    @BeforeAll
    public static void beforeAll(){
        RestAssured.baseURI ="http://localhost:8083/api/v1";
        request = given();
    }

    @Test
    @DisplayName("Get - AllCards - OK")
    public void allCards(){

        request
                .auth().oauth2(token)
                .when()
                .get("/accounts/")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Post - Card - Err")
    public void createAccount(){
        final String requestBody = "{\n" +
                "    \"card_number\": \"123344879\",\n" +
                "    \"expiration_date\": \"2022-11-11\",\n" +
                "    \"owner\": \"juli palma\",\n" +
                "    \"security_code\": \"483\",\n" +
                "    \"brand\": \"VISA\"\n" +
                "}";

        Response response = given()
                .auth().oauth2(token)
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/accounts/2/cards")
                .then()
                .extract().response();

        Assertions.assertEquals(403, response.statusCode());
    }

    @Test
    @DisplayName("Get - CardByID - OK")
    public void acountByID(){

        request
                .auth().oauth2(token)
                .when()
                .get("/accounts/1/cards/1")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Get - AllCards - OK")
    public void allCardsNoAuth(){

        request
                .when()
                .get("/accounts/")
                .then()
                .statusCode(200)
                .log().all();
    }

    @Test
    @DisplayName("Post - Card - Err")
    public void createAccountNoAuth(){
        final String requestBody = "{\n" +
                "    \"card_number\": \"123344879\",\n" +
                "    \"expiration_date\": \"2022-11-11\",\n" +
                "    \"owner\": \"juli palma\",\n" +
                "    \"security_code\": \"483\",\n" +
                "    \"brand\": \"VISA\"\n" +
                "}";

        Response response = given()
                .header("Content-type", "application/json")
                .and()
                .body(requestBody)
                .when()
                .post("/accounts/2/cards")
                .then()
                .extract().response();

        Assertions.assertEquals(404, response.statusCode());
    }

    @Test
    @DisplayName("Get - CardByID - Err")
    public void acountByIDNoAuth(){

        request
                .when()
                .get("/accounts/1/cards/1")
                .then()
                .statusCode(403)
                .log().all();
    }
}
