CREATE TABLE project_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES project_categories(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE freelancer_skills (
    freelancer_id INT NOT NULL REFERENCES freelancer_profiles(id) ON DELETE CASCADE,
    skill_id INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (freelancer_id, skill_id)
);