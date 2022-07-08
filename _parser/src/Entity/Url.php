<?php

declare(strict_types=1);

namespace App\Entity;

use App\Repository\UrlRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: UrlRepository::class)]
class Url
{
    public const REVIEW_TYPE = 1;

    #[ORM\Id]
    #[ORM\GeneratedValue(strategy: 'IDENTITY')]
    #[ORM\Column(type: 'integer')]
    private int $id;

    #[ORM\Column(type: 'string', length: 511)]
    private string $host;

    #[ORM\Column(type: 'text')]
    private string $path;

    #[ORM\Column(type: 'json', nullable: true)]
    private array $requestData = [];

    #[ORM\Column(type: 'boolean')]
    private ?bool $isResponseOK;

    #[ORM\Column(type: 'datetimetz_immutable')]
    private $lastRequest;

    #[ORM\Column(type: 'text', nullable: true)]
    private string $data;

    #[ORM\Column(type: 'smallint', nullable: true)]
    private int $type;

    #[ORM\Column(type: 'boolean', options: ['default' => false])]
    private ?bool $is_checked = false;

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

    public function getData(): ?string
    {
        return $this->data;
    }

    public function setData(?string $data): self
    {
        $this->data = $data;

        return $this;
    }

    public function getType(): ?int
    {
        return $this->type;
    }

    public function setType(?int $type): self
    {
        $this->type = $type;

        return $this;
    }

    public function isIsChecked(): ?bool
    {
        return $this->is_checked;
    }

    public function setIsChecked(bool $is_checked): self
    {
        $this->is_checked = $is_checked;

        return $this;
    }
}
