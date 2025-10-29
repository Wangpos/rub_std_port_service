import { body } from "express-validator";

export const createStudentValidator = [
  body("studentNumber").isString().notEmpty(),
  body("firstName").isString().notEmpty(),
  body("lastName").isString().notEmpty(),
  body("email").isEmail(),
  body("programId").isInt().toInt(),
  body("enrollmentYear").isInt().toInt(),
];

export const updateStudentValidator = [
  body("firstName").optional().isString(),
  body("lastName").optional().isString(),
  body("email").optional().isEmail(),
  body("programId").optional().isInt().toInt(),
  body("enrollmentYear").optional().isInt().toInt(),
];
