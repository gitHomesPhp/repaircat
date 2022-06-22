<?php

namespace App\Entity;

use App\Repository\UrlRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: UrlRepository::class)]
class Url
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column(type: 'integer')]
    private $id;

    #[ORM\Column(type: 'string', length: 511)]
    private $host;

    #[ORM\Column(type: 'text')]
    private $path;

    #[ORM\Column(type: 'json', nullable: true)]
    private $requestData = [];

    #[ORM\Column(type: 'boolean')]
    private $isResponseOK;

    #[ORM\Column(type: 'datetimetz_immutable')]
    private $lastRequest;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getHost(): ?string
    {
        return $this->host;
    }

    public function setHost(string $host): self
    {
        $this->host = $host;

        return $this;
    }

    public function getPath(): ?string
    {
        return $this->path;
    }

    public function setPath(string $path): self
    {
        $this->path = $path;

        return $this;
    }

    public function getRequestData(): ?array
    {
        return $this->requestData;
    }

    public function setRequestData(?array $requestData): self
    {
        $this->requestData = $requestData;

        return $this;
    }

    public function isIsResponseOK(): ?bool
    {
        return $this->isResponseOK;
    }

    public function setIsResponseOK(bool $isResponseOK): self
    {
        $this->isResponseOK = $isResponseOK;

        return $this;
    }

    public function getLastRequest(): ?\DateTimeImmutable
    {
        return $this->lastRequest;
    }

    public function setLastRequest(\DateTimeImmutable $lastRequest): self
    {
        $this->lastRequest = $lastRequest;

        return $this;
    }
}
