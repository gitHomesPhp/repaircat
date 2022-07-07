<?php

declare(strict_types=1);

namespace App\Entity;

use App\Repository\ScRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ScRepository::class)]
class Sc
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column(type: 'integer')]
    private $id;

    #[ORM\Column(type: 'bigint')]
    private $repaircat_id;

    #[ORM\Column(type: 'string', length: 511, nullable: true)]
    private $gis_id;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getRepaircatId(): ?string
    {
        return $this->repaircat_id;
    }

    public function setRepaircatId(string $repaircat_id): self
    {
        $this->repaircat_id = $repaircat_id;

        return $this;
    }

    public function getGisId(): ?string
    {
        return $this->gis_id;
    }

    public function setGisId(?string $gis_id): self
    {
        $this->gis_id = $gis_id;

        return $this;
    }
}
