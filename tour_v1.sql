CREATE TABLE "company_review"(
    "id" BIGINT NOT NULL,
    "rating" DECIMAL(8, 2) NOT NULL,
    "text" VARCHAR(255) NOT NULL,
    "traveler_id" BIGINT NOT NULL,
    "company_id" BIGINT NOT NULL
);
ALTER TABLE
    "company_review" ADD PRIMARY KEY("id");
CREATE TABLE "tour_moderator"(
    "id" BIGINT NOT NULL,
    "employee_id" BIGINT NOT NULL,
    "tour_id" BIGINT NOT NULL
);
CREATE INDEX "tour_moderator_employee_id_tour_id_index" ON
    "tour_moderator"("employee_id", "tour_id");
ALTER TABLE
    "tour_moderator" ADD PRIMARY KEY("id");
CREATE TABLE "super_user"(
    "id" BIGINT NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "email" BIGINT NOT NULL
);
ALTER TABLE
    "super_user" ADD PRIMARY KEY("id");
CREATE TABLE "tour_category"(
    "id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "tour_id" BIGINT NULL
);
ALTER TABLE
    "tour_category" ADD PRIMARY KEY("id");
CREATE TABLE "employee"(
    "id" BIGINT NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "company_id" BIGINT NOT NULL
);
ALTER TABLE
    "employee" ADD PRIMARY KEY("id");
CREATE TABLE "tour_group"(
    "id" BIGINT NOT NULL,
    "tour_id" BIGINT NOT NULL,
    "traveler_id" BIGINT NOT NULL
);
ALTER TABLE
    "tour_group" ADD PRIMARY KEY("id");
CREATE TABLE "company"(
    "id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NULL,
    "instagram" VARCHAR(255) NULL,
    "2gis" VARCHAR(255) NULL,
    "email" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "company" ADD PRIMARY KEY("id");
CREATE TABLE "tour"(
    "id" BIGINT NOT NULL,
    "company_id" BIGINT NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "start_time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "end_time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "group_size" BIGINT NOT NULL,
    "languages" VARCHAR(255) NOT NULL,
    "free_cancellation" BOOLEAN NOT NULL,
    "cancellation_condition" JSON NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "meeting_place" VARCHAR(255) NOT NULL,
    "arrival_place" VARCHAR(255) NOT NULL,
    "what_is_included" VARCHAR(255) NOT NULL,
    "what_to_prepare" VARCHAR(255) NOT NULL,
    "prohibitions" VARCHAR(255) NOT NULL,
    "price" BIGINT NOT NULL,
    "images" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "tour" ADD PRIMARY KEY("id");
CREATE TABLE "interests"(
    "id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "interests" ADD PRIMARY KEY("id");
CREATE TABLE "tour_review"(
    "id" BIGINT NOT NULL,
    "rating" DECIMAL(8, 2) NOT NULL,
    "text" VARCHAR(255) NOT NULL,
    "tour_id" BIGINT NOT NULL,
    "traveler_id" BIGINT NOT NULL
);
ALTER TABLE
    "tour_review" ADD PRIMARY KEY("id");
CREATE TABLE "traveler"(
    "id" BIGINT NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "instagram" VARCHAR(255) NULL,
    "date_of_birth" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "city" VARCHAR(255) NOT NULL,
    "country" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NULL,
    "interest" BIGINT NOT NULL,
    "favorite_tours" BIGINT NOT NULL
);

ALTER TABLE
    "traveler" ADD PRIMARY KEY("id");
ALTER TABLE
    "tour_moderator" ADD CONSTRAINT "tour_moderator_employee_id_foreign" FOREIGN KEY("employee_id") REFERENCES "employee"("id");
ALTER TABLE
    "traveler" ADD CONSTRAINT "traveler_interest_foreign" FOREIGN KEY("interest") REFERENCES "interests"("id");
ALTER TABLE
    "tour" ADD CONSTRAINT "tour_company_id_foreign" FOREIGN KEY("company_id") REFERENCES "company"("id");
ALTER TABLE
    "tour_group" ADD CONSTRAINT "tour_group_traveler_id_foreign" FOREIGN KEY("traveler_id") REFERENCES "traveler"("id");
ALTER TABLE
    "tour_moderator" ADD CONSTRAINT "tour_moderator_tour_id_foreign" FOREIGN KEY("tour_id") REFERENCES "tour"("id");
ALTER TABLE
    "tour_group" ADD CONSTRAINT "tour_group_tour_id_foreign" FOREIGN KEY("tour_id") REFERENCES "tour"("id");
ALTER TABLE
    "company_review" ADD CONSTRAINT "company_review_traveler_id_foreign" FOREIGN KEY("traveler_id") REFERENCES "traveler"("id");
ALTER TABLE
    "tour_review" ADD CONSTRAINT "tour_review_tour_id_foreign" FOREIGN KEY("tour_id") REFERENCES "tour"("id");
ALTER TABLE
    "company_review" ADD CONSTRAINT "company_review_company_id_foreign" FOREIGN KEY("company_id") REFERENCES "company"("id");
ALTER TABLE
    "employee" ADD CONSTRAINT "employee_company_id_foreign" FOREIGN KEY("company_id") REFERENCES "company"("id");
ALTER TABLE
    "tour_category" ADD CONSTRAINT "tour_category_tour_id_foreign" FOREIGN KEY("tour_id") REFERENCES "tour"("id");
ALTER TABLE
    "traveler" ADD CONSTRAINT "traveler_favorite_tours_foreign" FOREIGN KEY("favorite_tours") REFERENCES "tour"("id");
ALTER TABLE
    "tour_review" ADD CONSTRAINT "tour_review_traveler_id_foreign" FOREIGN KEY("traveler_id") REFERENCES "traveler"("id");