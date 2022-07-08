<?php

declare(strict_types=1);

namespace App\Service\GisParser\Review;

use App\Entity\Review;
use App\Entity\Sc;
use App\Entity\Url;
use Doctrine\ORM\EntityManagerInterface;
use stdClass;

class ReviewParser
{
    public function __construct(
        private EntityManagerInterface $em,
    ) {}

    public function handle(Url $url): void
    {
        $data = json_decode($url->getData());
        $scId = (int) explode('/', $url->getPath())[3];

        $this->parseData($data, $scId);
        $url->setIsChecked(true);
        $this->em->persist($url);
    }

    private function parseData(stdClass $dataObject, int $scId): void
    {
        $scRepo = $this->em->getRepository(Sc::class);
        $reviews = $dataObject->reviews;

        foreach ($reviews as $review) {
            $sc = $scRepo->findOneBy(['gis_id' => $scId]);

            $reviewEntity = new Review();
            $reviewEntity->setText($review->text);
            $reviewEntity->setRating($review->rating);
            $reviewEntity->setVisitorId(1);
            $reviewEntity->setSc($sc);

            $this->em->persist($reviewEntity);
        }
    }
}