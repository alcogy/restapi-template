import { Request, Response } from "express";
import { Sample, selectAllSamples, selectSampleByID, insertSample, updateSample, deleteSampleData } from "../models/sample";

export const getAllSamples = async (req: Request, res: Response) => {
  try {
    const result = await selectAllSamples();
    return res.status(200).send(result);
  } catch (e) {
    return res.status(500).send({"message": "ng"});
  }
};

export const getSampleByID = async (req: Request, res: Response) => {
  try {
    const id = req.params.id;
    const result = await selectSampleByID(id);
    return res.status(200).send(result);
  } catch (e) {
    return res.status(500).send({"message": "ng"});
  }
};

export const postSample = async (req: Request, res: Response) => {
  try {
    const id = await insertSample(req.body as Sample)
    return res.status(200).send({ id: id });
  } catch (e) {
    return res.status(500).send({"message": "ng"});
  }
};

export const putSample = async (req: Request, res: Response) => {
  try {
    await updateSample(req.body as Sample);
    return res.status(200).send({ message: "ok" });
  } catch (e) {
    return res.status(500).send({ message: "ng"});
  }
};

export const deleteSample = async (req: Request, res: Response) => {
  try {
    const id = req.params.id;
    await deleteSampleData(id);
    return res.status(200).send({ message: "ok"});
  } catch (e) {
    return res.status(500).send({message: "ng"});
  }
};