export interface User {
  id: number
  username: string
  firstname: string
  lastname: string
  email: string
}

export interface JWTPayload {
  MapClaims: {
    eat: number
    iat: number
    id: string
  }
  session: string
  searchFilters: string
}

export interface ICreateUser {
  username: string
  email: string
  firstname: string
  lastname: string
  password: string
}
export interface ILoginUser {
  username: string
  password: string
}

export interface IPerson {
  id: number
  age: number
  workclass: string
  fnlwgt: number
  education: string
  education_num: number
  marital_status: string
  occupation: string
  relationship: string
  race: string
  sex: string
  capital_gain: number
  capital_loss: number
  hours_per_week: number
  native_country: string
  income: string
}
