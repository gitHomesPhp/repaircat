<?php

declare(strict_types=1);

namespace App\Entity;

use App\Repository\ReviewRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ReviewRepository::class)]
class Review
{
    #[ORM\Id]
    #[ORM\GeneratedValue(strategy: 'IDENTITY')]
    #[ORM\Column(type: 'integer')]
    private int $id;

    #[ORM\ManyToOne(targetEntity: Sc::class)]
    #[ORM\JoinColumn(nullable: false)]
    private Sc $sc;

    #[ORM\Column(type: 'text')]
    private string $text;

    #[ORM\Column(type: 'smallint')]
    private int $rating;

    #[ORM\Column(type: 'bigint')]
    private int $visitor_id;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getSc(): ?Sc
    {
        return $this->sc;
    }

    public function setSc(?Sc $sc): self
    {
        $this->sc = $sc;

        return $this;
    }

    public function getText(): ?string
    {
        return $this->text;
    }

    public function setText(string $text): self
    {
        $this->text = $text;

        return $this;
    }

    public function getRating(): ?int
    {
        return $this->rating;
    }

    public function setRating(int $rating): self
    {
        $this->rating = $rating;

        return $this;
    }

    public function getVisitorId(): ?int
    {
        return $this->visitor_id;
    }

    public function setVisitorId(int $visitor_id): self
    {
        $this->visitor_id = $visitor_id;

        return $this;
    }
}
