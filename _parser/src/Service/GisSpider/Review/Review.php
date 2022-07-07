<?php

declare(strict_types=1);

namespace App\Service\GisSpider\Review;

class Review
{
    private const BASE_URL = 'https://public-api.reviews.2gis.com/2.0/branches';

    private static array $fields = [
        'meta.providers', 'meta.branch_rating', 'meta.branch_reviews_count',
        'meta.total_count', 'reviews.hiding_reason', 'reviews.is_verified',
    ];

    public function __construct(
        private readonly int $item_id,
        private string $url = '',
        private string $nextLink = '',
        private string $data = '',
    ) {
        $this->url = self::BASE_URL
            . "/{$this->item_id}/reviews?limit=12"
            . '&is_advertiser=true'
            . "&fields={$this->getFieldsUrlParams()}"
            . '&without_my_first_review=false'
            . '&rated=true'
            . '&sort_by=date_edited'
            . '&key=37c04fe6-a560-4549-b459-02309cf643ad'
            . '&locale=ru_RU'
        ;
    }

    public function getUrl(): string
    {
        return $this->url;
    }

    public function setNextLink(string $nextLink): self
    {
        $this->nextLink = $nextLink;

        return $this;
    }

    public function getNextLink(): string
    {
        return $this->nextLink;
    }

    private function getFieldsUrlParams(): string
    {
        return implode(',', self::$fields);
    }
}