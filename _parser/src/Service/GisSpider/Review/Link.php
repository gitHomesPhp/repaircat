<?php

declare(strict_types=1);

namespace App\Service\GisSpider\Review;


use function Symfony\Component\String\u;

class Link
{
    public const HOST = 'https://public-api.reviews.2gis.com';
    public const BASE_PATH = '/2.0/branches';

    private static array $fields = [
        'meta.providers', 'meta.branch_rating', 'meta.branch_reviews_count',
        'meta.total_count', 'reviews.hiding_reason', 'reviews.is_verified',
    ];

    private string $url;

    private int $itemId;

    public function getUrl(): string
    {
        return $this->url;
    }

    public function setUrl(string $url): self
    {
        $this->url = $url;

        return $this;
    }

    public function buildUrl(): self
    {
        $additionalPath = "/{$this->itemId}/reviews?limit=40"
            . '&is_advertiser=true'
            . "&fields={$this->getFieldsUrlParams()}"
            . '&without_my_first_review=false'
            . '&rated=true'
            . '&sort_by=date_edited'
            . '&key=37c04fe6-a560-4549-b459-02309cf643ad'
            . '&locale=ru_RU';

        $this->url = self::HOST . self::BASE_PATH . $additionalPath;

        return $this;
    }

    public function setItemId(int $itemId): self
    {
        $this->itemId = $itemId;

        return $this;
    }

    public function getPath(): string
    {
        $parsed = parse_url($this->url);

        return $parsed['path'] . '?' . $parsed['query'];
    }

    private function getFieldsUrlParams(): string
    {
        return implode(',', self::$fields);
    }
}