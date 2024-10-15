import cors from "cors";
import express from "express";
import bodyParser from "body-parser";
import dotenv from "dotenv";
import * as serv from "./services/sample";

const app = express();

dotenv.config();
const PORT = process.env.PORT ? process.env.PORT : 8080;

// CORS対策
app.use(cors({ origin: "*" }));

// データ取得
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// routing
app.get("/samples", serv.getAllSamples);
app.get("/sample/:id", serv.getSampleByID); 
app.post("/sample", serv.postSample); 
app.put("/sample", serv.putSample); 
app.delete("/sample/:id", serv.deleteSample); 

app.listen(PORT, () => { 
  console.log("Server running at PORT: ", PORT); 
}).on("error", (error) => {
  throw new Error(error.message);
});