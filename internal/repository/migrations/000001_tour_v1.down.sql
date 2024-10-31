ALTER TABLE
    "traveler" DROP CONSTRAINT "traveler_favorite_tours_foreign";

ALTER TABLE
    "tour_category" DROP CONSTRAINT "tour_category_tour_id_foreign";

ALTER TABLE
    "employee" DROP CONSTRAINT "employee_company_id_foreign";

ALTER TABLE
    "company_review" DROP CONSTRAINT "company_review_company_id_foreign";

ALTER TABLE
    "tour_review" DROP CONSTRAINT "tour_review_traveler_id_foreign";

ALTER TABLE
    "traveler" DROP CONSTRAINT "traveler_interest_foreign";

ALTER TABLE
    "tour_group" DROP CONSTRAINT "tour_group_tour_id_foreign";

ALTER TABLE
    "tour_moderator" DROP CONSTRAINT "tour_moderator_tour_id_foreign";

ALTER TABLE
    "tour_group" DROP CONSTRAINT "tour_group_traveler_id_foreign";

ALTER TABLE
    "tour" DROP CONSTRAINT "tour_company_id_foreign";

ALTER TABLE
    "tour_moderator" DROP CONSTRAINT "tour_moderator_employee_id_foreign";

ALTER TABLE
    "company_review" DROP CONSTRAINT "company_review_traveler_id_foreign";

ALTER TABLE
    "tour_review" DROP CONSTRAINT "tour_review_tour_id_foreign";

DROP TABLE "traveler";

DROP TABLE "tour_review";

DROP TABLE "interests";

DROP TABLE "tour";

DROP TABLE "company";

DROP TABLE "tour_group";

DROP TABLE "employee";

DROP TABLE "tour_category";

DROP TABLE "super_user";

DROP TABLE "tour_moderator";

DROP TABLE "company_review";