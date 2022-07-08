<?php

declare(strict_types=1);

namespace App\Entity;

use App\Repository\ScRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ScRepository::class)]
class Sc
{
    #[ORM\Id]
    #[ORM\GeneratedValue(strategy: 'IDENTITY')]
    #[ORM\Column(type: 'integer')]
    private int $id;

    #[ORM\Column(type: 'bigint')]
    private ?int $repaircat_id;

    #[ORM\Column(type: 'string', length: 511, nullable: true)]
    private ?string $gis_id;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getRepaircatId(): ?int
    {
        return $this->repaircat_id;
    }

    public function setRepaircatId(int $repaircat_id): self
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
