CREATE TYPE income AS ENUM (
    '>50K',
    '<=50K'
);

CREATE TYPE marital_status AS ENUM (
    'Married-civ-spouse',
    'Divorced',
    'Never-married',
    'Separated',
    'Widowed',
    'Married-spouse-absent',
    'Married-AF-spouse'
);

CREATE TYPE occupation AS ENUM (
    'Tech-support',
    'Craft-repair',
    'Other-service',
    'Sales',
    'Exec-managerial',
    'Prof-specialty',
    'Handlers-cleaners',
    'Machine-op-inspct',
    'Adm-clerical',
    'Farming-fishing',
    'Transport-moving',
    'Priv-house-serv',
    'Protective-serv',
    'Armed-Forces'
);

CREATE TYPE race AS ENUM (
    'White',
    'Asian-Pac-Islander',
    'Amer-Indian-Eskimo',
    'Other',
    'Black'
);

CREATE TYPE relationship AS ENUM (
    'Wife',
    'Own-child',
    'Husband',
    'Not-in-family',
    'Other-relative',
    'Unmarried'
);

CREATE TYPE sex AS ENUM (
    'Female',
    'Male'
);

CREATE TYPE workclass AS ENUM (
    'Private',
    'Self-emp-not-inc',
    'Self-emp-inc',
    'Federal-gov',
    'Local-gov',
    'State-gov',
    'Without-pay',
    'Never-worked'
);







