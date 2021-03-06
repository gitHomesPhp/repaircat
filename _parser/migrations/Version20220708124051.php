<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20220708124051 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE TABLE review (id SERIAL NOT NULL, sc_id INT NOT NULL, text TEXT NOT NULL, rating SMALLINT NOT NULL, visitor_id BIGINT NOT NULL, PRIMARY KEY(id))');
        $this->addSql('CREATE INDEX IDX_794381C6C8C329CD ON review (sc_id)');
        $this->addSql('CREATE TABLE sc (id SERIAL NOT NULL, repaircat_id BIGINT NOT NULL, gis_id VARCHAR(511) DEFAULT NULL, PRIMARY KEY(id))');
        $this->addSql('CREATE TABLE url (id SERIAL NOT NULL, host VARCHAR(511) NOT NULL, path TEXT NOT NULL, request_data JSON DEFAULT NULL, is_response_ok BOOLEAN NOT NULL, last_request TIMESTAMP(0) WITH TIME ZONE NOT NULL, data TEXT DEFAULT NULL, type SMALLINT DEFAULT NULL, is_checked BOOLEAN DEFAULT false NOT NULL, PRIMARY KEY(id))');
        $this->addSql('COMMENT ON COLUMN url.last_request IS \'(DC2Type:datetimetz_immutable)\'');
        $this->addSql('ALTER TABLE review ADD CONSTRAINT FK_794381C6C8C329CD FOREIGN KEY (sc_id) REFERENCES sc (id) NOT DEFERRABLE INITIALLY IMMEDIATE');
    }

    public function down(Schema $schema): void
    {
        // this down() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE SCHEMA public');
        $this->addSql('ALTER TABLE review DROP CONSTRAINT FK_794381C6C8C329CD');
        $this->addSql('DROP TABLE review');
        $this->addSql('DROP TABLE sc');
        $this->addSql('DROP TABLE url');
    }
}
