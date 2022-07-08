<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20220708011713 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE SEQUENCE url_id_seq INCREMENT BY 1 MINVALUE 1 START 1');
        $this->addSql('CREATE TABLE sc (id SERIAL NOT NULL, repaircat_id BIGINT NOT NULL, gis_id VARCHAR(511) DEFAULT NULL, PRIMARY KEY(id))');
        $this->addSql('CREATE TABLE url (id INT NOT NULL, host VARCHAR(511) NOT NULL, path TEXT NOT NULL, request_data JSON DEFAULT NULL, is_response_ok BOOLEAN NOT NULL, last_request TIMESTAMP(0) WITH TIME ZONE NOT NULL, data TEXT DEFAULT NULL, type SMALLINT DEFAULT NULL, PRIMARY KEY(id))');
        $this->addSql('COMMENT ON COLUMN url.last_request IS \'(DC2Type:datetimetz_immutable)\'');
    }

    public function down(Schema $schema): void
    {
        // this down() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE SCHEMA public');
        $this->addSql('DROP SEQUENCE url_id_seq CASCADE');
        $this->addSql('DROP TABLE sc');
        $this->addSql('DROP TABLE url');
    }
}
