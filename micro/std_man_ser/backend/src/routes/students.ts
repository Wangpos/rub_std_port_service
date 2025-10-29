import { Router } from "express";
import * as studentController from "../controllers/studentController";
import {
  createStudentValidator,
  updateStudentValidator,
} from "../validators/studentValidators";
import validate from "../middleware/validate";

const router = Router();

router.get("/", studentController.listStudents);
router.post(
  "/",
  createStudentValidator,
  validate,
  studentController.createStudent
);
router.get("/:id", studentController.getStudent);
router.put(
  "/:id",
  updateStudentValidator,
  validate,
  studentController.updateStudent
);
router.delete("/:id", studentController.deleteStudent);

export default router;
